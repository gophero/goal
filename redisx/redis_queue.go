package redisx

import (
	"context"
	"fmt"
	"github.com/gophero/goal/queue"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

const (
	maxRetry = 6
)

// ==============================
// normal queue
// ==============================

func NewQueue(rc Client, key string) queue.Queue {
	return &queueImpl{rc: rc, key: key}
}

type queueImpl struct {
	rc  Client // redis client
	key string // list key
}

func (q *queueImpl) Push(v ...any) error {
	return q.rc.LPush(context.Background(), q.key, v...).Err()
}

func (q *queueImpl) Pop() (string, error) {
	cmd := q.rc.RPop(context.Background(), q.key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Val(), nil
}

func (q *queueImpl) PopN(n int) ([]string, error) {
	cmd := q.rc.RPopCount(context.Background(), q.key, n)
	if cmd.Err() != nil {
		return []string{}, cmd.Err()
	}
	return cmd.Val(), nil
}

func (q *queueImpl) BPop() (string, error) {
	cmd := q.rc.BRPop(context.Background(), 0, q.key) // 没有数据时阻塞
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	if len(cmd.Val()) > 0 { // []string 为两个元素，第一个为key，第二个为value
		return cmd.Val()[1], nil
	}
	return "", nil
}

func (q *queueImpl) Len() uint64 {
	return uint64(q.rc.LLen(context.Background(), q.key).Val())
}

// ==============================
// unique queue
// ==============================

func NewUniqueQueue(rc Client, key string) queue.Queue {
	// 使用hash标签模式将key的hash slot映射到集群的同一个节点，避免 CROSSSLOT 错误，见：
	// https://hackernoon.com/resolving-the-crossslot-keys-error-with-redis-cluster-mode-enabled
	// https://redis.io/docs/reference/cluster-spec/
	// 因为queue和set会用到事务，事务操作的key可能在集群环境下被映射到不同的hash slot，从而出错：
	// CROSSSLOT Keys in request don't hash to the same slot
	// 给key加上标签"{}"后，只有"{}"中的部分会用来计算hash，此时queue和set的key就被hash到同一个slot下
	key = "{" + key + "}"
	setkey := key + "_set"
	return &uniqueQueue{
		queueImpl: NewQueue(rc, key).(*queueImpl),
		setkey:    setkey,
		set:       NewSet(rc, setkey),
	}
}

func createSetWhenExists(rc Client, key string) (string, Set) {
	var setExists bool
	var setkey string
	var setkeyPrefix = key + "_"
	if cmd := rc.Keys(context.Background(), setkeyPrefix+"*"); cmd.Err() != nil {
		panic(cmd.Err())
	} else {
		setExists = len(cmd.Val()) > 0
		if setExists {
			setkey = cmd.Val()[len(cmd.Val())-1]
		} else {
			setkey = fmt.Sprintf("%s%d", setkeyPrefix, time.Now().Unix())
		}
	}
	return setkey, NewSet(rc, setkey)
}

type uniqueQueue struct {
	mu sync.Mutex
	*queueImpl

	set    Set
	setkey string
}

func (uq *uniqueQueue) retryPush(vs ...any) error {
	var ctx = context.TODO()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("error recovered: %v\n", err)
			}
		}()
		times := 0
		for {
			cmd := uq.rc.LPush(ctx, uq.key, vs...)
			if cmd.Err() == nil {
				break
			}
			times++
			if times > maxRetry {
				fmt.Println(fmt.Errorf("ERROR: push failed after retry the max %d times", maxRetry))
				break
			}
			fmt.Printf("ERROR: push failed, will retry at the %d time after %d second\n", times, 1<<times)
			time.Sleep(time.Second * (1 << times))
		}
	}()
	return nil
}

func (uq *uniqueQueue) Push(vs ...any) error {
	uq.mu.Lock()
	defer uq.mu.Unlock()

	// 去重
	bs := uq.set.IsMembers(vs...)
	var dest []any
	for i, v := range vs {
		if !bs[i] {
			dest = append(dest, v)
		}
	}
	if len(dest) == 0 {
		return nil
	}

	// 使用Pipliner开启事务保证一致性
	var ctx = context.TODO()
	_, err := uq.queueImpl.rc.TxPipelined(ctx, func(pip redis.Pipeliner) error {
		cmd := pip.LPush(ctx, uq.key, dest...)
		if cmd.Err() != nil {
			return cmd.Err()
		}
		return pip.SAdd(context.Background(), uq.setkey, dest...).Err()
	})
	return err
}

func (uq *uniqueQueue) Pop() (string, error) {
	if s, err := uq.queueImpl.Pop(); err != nil {
		return "", err
	} else {
		// 出错重新放入队列
		if err := uq.set.Rem(s); err != nil {
			fmt.Printf("rem [%v] from set error, will retry %d times to push to queue: %v\n", s, maxRetry, err)
			return s, uq.retryPush(s)
			//return s, err
		}
		return s, nil
	}
}

func (uq *uniqueQueue) PopN(n int) ([]string, error) {
	//var vs []string
	//var ctx = context.TODO()
	//_, err := uq.queueImpl.rc.TxPipelined(ctx, func(pip redis.Pipeliner) error {
	//	sscmd := pip.RPopCount(ctx, uq.key, n)
	//	if sscmd.Err() != nil {
	//		return sscmd.Err()
	//	}
	//	vs = sscmd.Val()
	//	if len(vs) == 0 {
	//		return nil
	//	}
	//	var ss = make([]any, len(vs))
	//	for i, v := range vs {
	//		ss[i] = v
	//	}
	//	return pip.SRem(ctx, uq.setkey, ss...).Err()
	//})
	//return vs, err
	if vs, err := uq.queueImpl.PopN(n); err != nil {
		return vs, err
	} else {
		if len(vs) == 0 {
			return []string{}, nil
		}
		var ss = make([]any, len(vs))
		for i, v := range vs {
			ss[i] = v
		}
		if err := uq.set.Rem(ss...); err != nil {
			//fmt.Printf("rem [%v] from set error, will retry %d times to push to queue: %v\n", ss, maxRetry, err)
			return vs, uq.retryPush(ss...)
		}
		return vs, nil
	}
}

func (uq *uniqueQueue) BPop() (string, error) {
	if s, err := uq.queueImpl.BPop(); err != nil { // 阻塞
		return "", err
	} else {
		// 删除set中的元素失败，重新放入队列
		if err := uq.set.Rem(s); err != nil {
			fmt.Printf("rem [%s] from set error, will retry %d times to push to queue: %v\n", s, maxRetry, err)
			return s, uq.retryPush(s)
			//return s, err
		}
		return s, err
	}
}

// ==============================
// bounded queue
// ==============================

func NewBoundedQueue(rc Client, key string, cap uint64) queue.BoundedQueue {
	return &boundedQueue{
		queueImpl: NewQueue(rc, key).(*queueImpl),
		cap:       cap,
	}
}

type boundedQueue struct {
	*queueImpl

	cap uint64
}

func (bq *boundedQueue) Cap() uint64 {
	return bq.cap
}

func (bq *boundedQueue) Push(vs ...any) error {
	if bq.Len() >= bq.cap {
		return fmt.Errorf("quene reachs max capacity, can not push now")
	}
	return bq.queueImpl.Push(vs...)
}

// ==============================
// bounded unique queue
// ==============================

func NewBoundedUniqueQueue(rc Client, key string, cap uint64) queue.BoundedQueue {
	return &boundedUniqueQueue{
		uniqueQueue: NewUniqueQueue(rc, key).(*uniqueQueue),
		cap:         cap,
	}
}

type boundedUniqueQueue struct {
	*uniqueQueue

	cap uint64
}

func (bq *boundedUniqueQueue) Cap() uint64 {
	return bq.cap
}

func (bq *boundedUniqueQueue) Push(vs ...any) error {
	if bq.Len() >= bq.cap {
		return fmt.Errorf("quene reachs max capacity, can not push now")
	}
	return bq.uniqueQueue.Push(vs...)
}

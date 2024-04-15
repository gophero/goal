package redisx_test

import (
	"fmt"
	"github.com/gophero/goal/redisx"
	"github.com/gophero/goal/testx"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("begin to test redis queue")
	code := m.Run()
	fmt.Println("test redis queue end")
	os.Exit(code)
}

func TestNormalQueue(t *testing.T) {
	var err error
	rc := testx.NewMiniRedis().(*redis.Client)
	q := redisx.NewQueue(rc, "test_normal_queue")
	err = q.Push("key1", "key2")
	assert.True(t, q.Len() == 2)
	i := 0
	for {
		s, err1 := q.Pop()
		if err1 == redis.Nil {
			break
		}
		assert.True(t, err1 == nil)
		assert.True(t, s != "")
		i++
	}
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	assert.True(t, i == 2)
	assert.True(t, q.Len() == 0)
}

func TestUniqueQueue(t *testing.T) {
	var err error
	rc := testx.NewMiniRedis().(*redis.Client)
	q := redisx.NewUniqueQueue(rc, "test_uniqueue")
	err = q.Push("key1", "key2")
	assert.True(t, q.Len() == 2)
	i := 0
	for {
		s, err1 := q.Pop()
		if err1 == redis.Nil {
			break
		}
		assert.True(t, s != "")
		i++
	}
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	assert.True(t, i == 2)
	assert.True(t, q.Len() == 0)
}

func TestUniqueQueue_Push(t *testing.T) {
	var err error
	rc := testx.NewMiniRedis().(*redis.Client)
	q := redisx.NewUniqueQueue(rc, "test_uniqueue")
	err = q.Push("key1", "key2")
	assert.True(t, q.Len() == 2)
	err = q.Push("key1", "key2") // 重复的数据，无效
	assert.True(t, q.Len() == 2)
	i := 0
	for {
		s, err1 := q.Pop()
		if err1 == redis.Nil {
			break
		}
		assert.True(t, s != "")
		i++
	}
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	assert.True(t, i == 2)
	assert.True(t, q.Len() == 0)
}

func TestUniqueQueue_PopN(t *testing.T) {
	var err error
	rc := testx.NewMiniRedis().(*redis.Client)
	q := redisx.NewUniqueQueue(rc, "test_uniqueue")
	err = q.Push("key1", "key2")
	assert.True(t, q.Len() == 2)
	s, err := q.PopN(2)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	assert.True(t, q.Len() == 0)
	assert.True(t, len(s) == 2)
}

type uniqueue struct {
	redisx.UniqueQueue
}

func (uq *uniqueue) retryPush(ch chan struct{}, vs ...any) error {
	//var ctx = context.TODO()
	go func() {
		times := 0
		for {
			//cmd := uq.rc.LPush(ctx, uq.key, vs...)
			//if cmd.Err() == nil {
			//	break
			//}
			times++
			if times > redisx.MaxRetry {
				fmt.Println(fmt.Errorf("ERROR: push failed after retry the max %d times", redisx.MaxRetry))
				ch <- struct{}{}
				break
			}
			fmt.Printf("ERROR: push failed, will retry at the %d time after %d second\n", times, 1<<times)
			time.Sleep(time.Second * (1 << times))
		}
	}()
	return nil
}

func TestUniqueQueue_RetryPush(t *testing.T) {
	ch := make(chan struct{})
	rc := testx.NewMiniRedis().(*redis.Client)
	q := uniqueue{UniqueQueue: redisx.NewExportUniqueQueue(rc, "test_unique")}
	err := q.retryPush(ch, "key1", "key2")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	fmt.Println("retry push is asynchronously")
	<-ch
}

func TestBoundedQueue(t *testing.T) {
	rc := testx.NewMiniRedis().(*redis.Client)
	q := redisx.NewBoundedQueue(rc, "test_bounded_queue", 2)
	err := q.Push("key1", "key2")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	err = q.Push("key1", "key2") // 容量已满，再push会出错
	assert.True(t, err != nil)
	s, err := q.Pop()
	assert.True(t, err == nil)
	assert.True(t, s == "key1")
	err = q.Push("key1")
	assert.True(t, err == nil) // 又可以继续添加了
}

func TestBoundedUniqueQueue(t *testing.T) {
	rc := testx.NewMiniRedis().(*redis.Client)
	q := redisx.NewBoundedUniqueQueue(rc, "test_unique_bounded_queue", 2)
	err := q.Push("key1", "key2")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	assert.True(t, q.Len() == 2)
	err = q.Push("key1", "key2") // 容量已满，再push会出错
	assert.True(t, err != nil)
	assert.True(t, q.Len() == 2)
	s, err := q.Pop()
	assert.True(t, err == nil)
	assert.True(t, s == "key1")
	assert.True(t, q.Len() == 1)
	err = q.Push("key2") // 重复了，无法成功添加
	assert.True(t, err == nil)
	assert.True(t, q.Len() == 1)
	err = q.Push("key1")
	assert.True(t, err == nil)
	assert.True(t, q.Len() == 2)
}

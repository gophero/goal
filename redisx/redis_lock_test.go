package redisx_test

import (
	"context"
	"fmt"
	"github.com/gophero/goal/redisx"
	"github.com/gophero/goal/testx"
	"testing"
	"time"

	"github.com/gophero/goal/assert"
	"github.com/redis/go-redis/v9"
)

func conn() redis.Cmdable {
	return NewMiniRedis()
}

func testConn(tl *testx.Logger) {
	tl.Case("test connect redis")
	c := NewMiniRedis()
	assert.NoneNil(c)
}

func testLock(tl *testx.Logger) {
	tl.Case("test reentry lock")
	var exp uint32 = 2 // 超时时间
	lock := redisx.NewLock(conn(), "test-key", exp)
	_, err := lock.Release() // 先释放锁
	tl.Require(err == nil, "release error: %v", err)
	b, err := lock.Acquire() // 加锁
	tl.Require(err == nil, "acquire error: %v", err)
	tl.Require(b, "acquire should be success")

	time.Sleep(time.Second)
	b, err = lock.Acquire() // 再加锁，重入，更新过期时间
	tl.Require(err == nil, "acquire error: %v", err)
	tl.Require(b, "acquire should be success")

	time.Sleep(time.Second)
	b, err = lock.Acquire() // 再加锁，继续重入
	tl.Require(err == nil, "acquire error: %v", err)
	tl.Require(b, "acquire should be failed after timeout")

	tl.Case("test lock timeout and reentry")
	time.Sleep(time.Second * time.Duration(exp)) // 超时时间后
	b, err = lock.Acquire()                      // 再加锁
	tl.Require(err == nil, "acquire error: %v", err)
	tl.Require(b, "acquire should be success")
}

func testLock1(tl *testx.Logger) {
	tl.Case("test reentry lock many times")
	n := 1000
	var exp uint32 = 2 // 超时时间
	lock := redisx.NewLock(conn(), "test-key1", exp)
	for i := 0; i < n; i++ {
		b, err := lock.Acquire() // 重入加锁
		tl.Require(err == nil, "acquire error: %v", err)
		tl.Require(b, "acquire many time should be success when not timeout")
	}
}

func testLock2(tl *testx.Logger) {
	tl.Case("test create many different lock and use them")

	client := conn()

	var exp uint32 = 10 // 超时时间
	n := 1000
	for i := 0; i < n; i++ {
		lock := redisx.NewLock(client, "test-key1", exp)
		b, err := lock.Acquire() // 重入加锁
		tl.Require(err == nil, "acquire error: %v", err)
		if i == 0 {
			tl.Require(b, "the first lock acquires should be success")
		} else {
			tl.Require(!b, "not the first lock with the same key acquires should be failed")
		}
	}
}

func testLock3(tl *testx.Logger) {
	tl.Case("test acquire and release lock")
	n := 1000
	var exp uint32 = 2 // 超时时间
	lock := redisx.NewLock(conn(), "test-key1", exp)
	for i := 0; i < n; i++ {
		b, err := lock.Acquire() // 重入加锁
		tl.Require(err == nil, "acquire error: %v", err)
		tl.Require(b, "acquire should be success")
		b, err = lock.Release()
		tl.Require(err == nil, "acquire error: %v", err)
		tl.Require(b, "release should be success")
	}
}

func TestRedixLock(t *testing.T) {
	tl := testx.Wrap(t)
	tl.Case("test acquire lock")
	testConn(tl)
	testLock(tl)
	testLock1(tl)
	testLock2(tl)
	testLock3(tl)
}

func TestRedisGet(t *testing.T) {
	client := testx.NewRedisCluster()
	if cc, ok := client.(*redis.ClusterClient); ok {
		ctx := context.Background()
		err := cc.ForEachMaster(ctx, func(ctx context.Context, c *redis.Client) error {
			key := cc.Get(ctx, "{1679091c5a880faf6fb5e6087eb1b2dc}*")
			fmt.Printf("keys: %v \n", key)
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
}

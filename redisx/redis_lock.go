package redisx

import (
	"context"
	"fmt"
	"github.com/gophero/goal/stringx"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync/atomic"
)

const (
	randomLen = 16  // 随机id长度
	tolerance = 500 // 冗余时间（ms）
)

var (
	// 可重入LUA加锁脚本，重入后更新过期时间
	lockScript = redis.NewScript(`if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`)

	// 释放锁脚本
	delScript = redis.NewScript(`if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`)
)

// A Lock is a redis lock.
type Lock struct {
	store   redis.Cmdable
	seconds uint32
	key     string
	id      string
}

// NewLock returns a Lock.
func NewLock(store redis.Cmdable, key string, expireSec uint32) *Lock {
	return &Lock{
		store:   store,
		key:     key,
		id:      stringx.Randn(randomLen),
		seconds: expireSec,
	}
}

// Acquire acquires the lock.
func (rl *Lock) Acquire() (bool, error) {
	return rl.AcquireCtx(context.Background())
}

// AcquireCtx acquires the lock with the given ctx.
func (rl *Lock) AcquireCtx(ctx context.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)
	cmd := lockScript.Run(ctx, rl.store, []string{rl.key}, []string{
		rl.id, strconv.Itoa(int(seconds)*1000 + tolerance),
	})
	resp := cmd.Val()
	err := cmd.Err()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("error on acquiring lock for %s, error: %w", rl.key, err)
	} else if resp == nil {
		return false, nil
	}

	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return true, nil
	}
	return false, fmt.Errorf("unknown reply when acquiring lock for %s, resp: %v", rl.key, resp)
}

// Release releases the lock.
func (rl *Lock) Release() (bool, error) {
	return rl.ReleaseCtx(context.Background())
}

// ReleaseCtx releases the lock with the given ctx.
func (rl *Lock) ReleaseCtx(ctx context.Context) (bool, error) {
	cmd := delScript.Run(ctx, rl.store, []string{rl.key}, []string{rl.id})
	resp := cmd.Val()
	err := cmd.Err()
	if err != nil {
		return false, err
	}

	reply, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return reply == 1, nil
}

// SetExpire sets the expiration.
func (rl *Lock) SetExpire(seconds int) {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
}

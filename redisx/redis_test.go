package redisx_test

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/gophero/goal/redisx"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewMiniRedis() redisx.Client {
	// 测试用miniredis
	mr, err := miniredis.Run()
	if err != nil {
		panic(fmt.Errorf("new test redis error: %v", err))
	}
	// 使用miniredis创建client
	client := redis.NewClient(&redis.Options{
		Addr:         mr.Addr(),
		Password:     "",
		DB:           0,
		MaxRetries:   5,
		MinIdleConns: 2,
		TLSConfig:    nil,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis error: %s", err.Error()))
	}
	fmt.Printf("redis connected, url: %s\n", client.Conn().String())
	return client
}

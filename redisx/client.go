package redisx

import "github.com/redis/go-redis/v9"

type Client interface {
	redis.UniversalClient
}

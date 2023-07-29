package cache

import (
	"github.com/go-redis/redis/v8"
)

var _ MessageCache = (*customMessageCache)(nil)

type (
	// MessageCache is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageCache.
	MessageCache interface {
	}
	customMessageCache struct {
		*redis.Client
	}
)

func NewMessageCache() MessageCache {
	return &customMessageCache{
		Client: RedisClient,
	}
}

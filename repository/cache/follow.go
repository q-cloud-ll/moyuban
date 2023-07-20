package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var _ FollowCache = (*customFollowCache)(nil)

type (
	  // FollowCache is an interface to be customized, add more methods here,
	  // and implement the added methods in customFollowCache.
	  FollowCache interface {
	  }
	  customFollowCache struct {
		  *redis.Client
	  }
)

func NewFollowCache() FollowCache {
	  return &customFollowCache{
		  Client: RedisClient,
	  }
}


package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var _ UserCache = (*customUserCache)(nil)

type (
	// UserCache is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCache.
	UserCache interface {
		SavePhoneMsg(ctx context.Context, phone, code string) error
		GetPhoneMsg(ctx context.Context, phone string) (code string, err error)
	}
	customUserCache struct {
		*redis.Client
	}
)

func NewUserCache() UserCache {
	return &customUserCache{
		Client: RedisClient,
	}
}
func (c *customUserCache) SavePhoneMsg(ctx context.Context, phone, code string) error {
	return c.Set(ctx, getRedisKey(phone), code, 2*time.Minute).Err()
}

func (c *customUserCache) GetPhoneMsg(ctx context.Context, phone string) (code string, err error) {
	return c.Get(ctx, getRedisKey(phone)).Result()
}

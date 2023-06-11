package redis

import (
	"context"
	"os"
	"project/setting"
	"sync"

	"github.com/go-redis/redis/v8"

	"go.uber.org/zap"
)

var (
	once   sync.Once
	client *redis.Client
)

func Init(redisCfg *setting.RedisConfig) error {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Host,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
		pong, err := client.Ping(context.Background()).Result()
		if err != nil {
			zap.L().Error("redis connect ping failed, err:", zap.Error(err))
			os.Exit(0)
			return
		} else {
			zap.L().Info("redis connect ping response:", zap.String("pong", pong))
		}
	})

	return nil
}

func Close() {
	_ = client.Close()
}
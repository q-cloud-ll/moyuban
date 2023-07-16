package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"math"
	"project/consts"
	"project/types"
	"time"
)

var _ StarCache = (*customStarCache)(nil)

type (
	// StarCache is an interface to be customized, add more methods here,
	// and implement the added methods in customStarCache.
	StarCache interface {
		StarPost(ctx context.Context, uid, pid string, dir float64) error
	}
	customStarCache struct {
		*redis.Client
	}
)

func NewStarCache() StarCache {
	return &customStarCache{
		Client: RedisClient,
	}
}

func (c *customStarCache) StarPost(ctx context.Context, uid, pid string, value float64) error {
	//TODO implement me

	// 先查当前用户给当前帖子的投票记录
	ov := c.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+pid), uid).Val()

	if value == ov {
		return consts.StarPostRepeatedErr
	}

	pipeline := c.Pipeline()
	// 建立一个所有帖子点赞的集合
	pipeline.SAdd(ctx, getRedisKey(KeyPostLikedSetPF), pid)
	// 建立当前帖子所点赞用户的集合
	pipeline.SAdd(ctx, getRedisKey(pid), uid)
	_, err := pipeline.Exec(ctx)
	if err != nil {
		return err
	}

	// calculate post score
	var psd types.StarPostDetail
	var op float64
	if value > op {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline1 := c.TxPipeline()
	// 记录用户行为
	if pipeline1.Exists(ctx, getStarRedisKey(pid, uid)).Val() < 1 {
		pipeline1.HMSet(ctx, getStarRedisKey(pid, uid),
			"post_id", pid,
			"user_id", uid,
			"status", value,
			"created_at", time.Now().Format("2006-01-02 15:04:05"),
			"updated_at", 0)
	} else {
		pipeline1.HSet(ctx, getStarRedisKey(pid, uid),
			psd.UpdatedAt, time.Now().Format("2006-01-02 15:04:05"))
	}
	var dif int64
	if value == 1 {
		dif = 1
	} else if (value == 0 && ov != -1) || (value == -1 && ov == 1) {
		dif = -1
	} else {
		dif = 0
	}

	if pipeline1.HExists(ctx, getRedisKey(KeyPostLikedCounterHSetPF), pid).Val() {
		pipeline1.HSet(ctx, getRedisKey(KeyPostLikedCounterHSetPF), pid, dif)
	} else {
		pipeline1.HIncrBy(ctx, getRedisKey(KeyPostLikedCounterHSetPF), pid, dif)
	}

	// 计算zset排行分值
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, pid)
	if value == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+pid), pid)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+pid), &redis.Z{
			Score:  value,
			Member: uid,
		})
	}
	_, err = pipeline.Exec(ctx)
	return err
}

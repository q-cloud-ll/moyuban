package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"project/consts"
	"project/types"
	"time"
)

var _ PostCache = (*customPostCache)(nil)

type (
	PostCache interface {
		CreatePost(ctx context.Context, pid, cid string) error
		GetPostIdsOrder(ctx context.Context, p *types.PostListReq) ([]string, error)
		GetCommunityPostsIdsInorder(ctx context.Context, p *types.PostListReq) ([]string, error)
	}
	customPostCache struct {
		*redis.Client
	}
)

func NewPostCache() PostCache {
	return &customPostCache{
		Client: RedisClient,
	}
}

// GetCommunityPostsIdsInorder 根据社区查询帖子id
func (c *customPostCache) GetCommunityPostsIdsInorder(ctx context.Context, p *types.PostListReq) ([]string, error) {
	//TODO implement me
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == consts.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	cKey := getRedisKey(KeyCommunitySetPF + p.CommunityID)
	key := orderKey + p.CommunityID
	if c.Exists(ctx, key).Val() < 1 {
		pipeline := c.Pipeline()
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Aggregate: "MAX",
			Keys:      []string{cKey, orderKey},
		})
		pipeline.Expire(ctx, key, 3*time.Second)
		_, err := pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}

	return c.GetIdsFormKey(ctx, key, p.Page, p.PageSize)

}

// GetPostIdsOrder 从redis中获取id
func (c *customPostCache) GetPostIdsOrder(ctx context.Context, p *types.PostListReq) ([]string, error) {
	//TODO implement me
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == consts.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	return c.GetIdsFormKey(ctx, key, p.Page, p.PageSize)
}

// GetIdsFormKey 根据时间或者得分，获取帖子id的排行
func (c *customPostCache) GetIdsFormKey(ctx context.Context, key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1

	return c.ZRevRange(ctx, key, start, end).Result()
}

// CreatePost 创建帖子
func (c *customPostCache) CreatePost(ctx context.Context, pid, cid string) error {
	//TODO implement me
	pipeline := c.TxPipeline()
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pid,
	})

	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pid,
	})

	// 将帖子id添加到社区的set 例如forum:community:3826223906033666 1944204179673088
	cKey := getRedisKey(KeyCommunitySetPF + cid)

	pipeline.SAdd(ctx, cKey, pid)
	_, err := pipeline.Exec(ctx)

	return err
}

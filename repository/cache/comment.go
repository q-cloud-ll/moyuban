package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mlogclub/simple/common/strs"
	"project/consts"
	"project/types"
	"strconv"
	"time"
)

var _ CommentCache = (*customCommentCache)(nil)

type (
	// CommentCache is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentCache.
	CommentCache interface {
		CreateComment(ctx context.Context, commentId, postId, pid int64) error
		IncrChildrenNum(ctx context.Context, commentId int64) (err error)
		GetCommentsIdsInOrder(ctx context.Context, req *types.CommentListReq) ([]string, error)
		GetChildrenNum(ctx context.Context, commentId int64) (int64, error)
		GetCommentStar(ctx context.Context, commentId int64) (int64, error)
	}

	customCommentCache struct {
		*redis.Client
	}
)

func NewCommentCache() CommentCache {
	return &customCommentCache{
		Client: RedisClient,
	}
}

func (c *customCommentCache) GetCommentStar(ctx context.Context, commentId int64) (int64, error) {
	//TODO implement me
	num, err := c.HGet(ctx, getRedisKey(KeyCommentLikedCounterHSetPF), strconv.FormatInt(commentId, 10)).Result()
	if strs.IsBlank(num) {
		return 0, nil
	}
	likeNum, _ := strconv.ParseInt(num, 10, 64)

	return likeNum, err
}
func (c *customCommentCache) GetChildrenNum(ctx context.Context, commentId int64) (int64, error) {
	num, err := c.Get(ctx, getRedisKey(KeyCommentChildrenNumSetPF+strconv.FormatInt(commentId, 10))).Result()
	if err != nil {
		return 0, err
	}
	res, _ := strconv.ParseInt(num, 10, 64)

	return res, err
}

func (c *customCommentCache) GetCommentsIdsInOrder(ctx context.Context, req *types.CommentListReq) ([]string, error) {
	key := getRedisKey(KeyCommentTimeZSetPF + req.PostId)
	if req.Order == consts.OrderScore {
		key = getRedisKey(KeyCommentTimeZSetPF + req.PostId)
	}

	return NewPostCache().GetIdsFormKey(ctx, key, req.Page, req.PageSize)
}

func (c *customCommentCache) CreateComment(ctx context.Context, commentId, postId, pid int64) error {
	var pidStr, postIdStr string
	postIdStr = strconv.FormatInt(postId, 10)
	if pid == 0 {
		pidStr = ""
	} else {
		pidStr = strconv.FormatInt(pid, 10)
	}
	pipeline := c.TxPipeline()

	pipeline.ZAdd(ctx, getRedisKey(KeyCommentTimeZSetPF+postIdStr+pidStr), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: commentId,
	})
	pipeline.ZAdd(ctx, getRedisKey(KeyCommentScoreZSetPF+postIdStr+pidStr), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: commentId,
	})
	_, err := pipeline.Exec(ctx)

	return err
}

func (c *customCommentCache) IncrChildrenNum(ctx context.Context, commentId int64) (err error) {
	//TODO implement me
	if c.Exists(ctx, getRedisKey(KeyCommentChildrenNumSetPF+strconv.FormatInt(commentId, 10))).Val() < 1 {
		err = c.Set(ctx, getRedisKey(KeyCommentChildrenNumSetPF+strconv.FormatInt(commentId, 10)), 1, 0).Err()
		return
	}
	err = c.Incr(ctx, getRedisKey(KeyCommentChildrenNumSetPF+strconv.FormatInt(commentId, 10))).Err()

	return
}

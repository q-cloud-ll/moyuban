package cache

import (
	"context"
	"math"
	"project/consts"
	"project/repository/db/model"
	"project/types"
	"project/utils"
	"project/utils/snowflake"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ StarCache = (*customStarCache)(nil)

type (
	// StarCache is an interface to be customized, add more methods here,
	// and implement the added methods in customStarCache.
	StarCache interface {
		StarPost(ctx context.Context, uid, pid string, dir float64) error
		UpdateStarDetailFromRedisToMySQL(ctx context.Context, db *gorm.DB) (err error)
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

/**
redis点赞逻辑key:
set: getRedisKey(KeyPostLikedSetPF) 存放所有被点赞帖子集合
set: getRedisKey(pid) 存放某个帖子下的所有点赞用户
hash: getStarRedisKey(pid, uid) 记录用户行为的hash，刷表以后可以分析用户行为
hash: getRedisKey(KeyPostLikedCounterHSetPF) 维护一个帖子的点赞计数器
*/

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

func (c *customStarCache) UpdateStarDetailFromRedisToMySQL(ctx context.Context, db *gorm.DB) (err error) {
	if db == nil {
		return consts.DBErrNotFound
	}
	zap.L().Info("UpdateStarDetailFromRedisToMySQL 定时任务执行")
	var f types.StarPostDetail
	for c.SCard(ctx, getRedisKey(KeyPostLikedSetPF)).Val() != 0 {
		postId := c.SPop(ctx, getRedisKey(KeyPostLikedSetPF)).Val()
		for c.SCard(ctx, getRedisKey(postId)).Val() != 0 {
			userId := c.SPop(ctx, getRedisKey(postId)).Val()
			err = c.HMGet(ctx, getStarRedisKey(postId, userId),
				"user_id", "post_id", "status", "create_at", "update_at").Scan(&f)
			if err != nil {
				zap.L().Error("UpdateStarDetailFromRedisToMySQL data failed", zap.Error(err))
				continue
			}
			status, _ := strconv.Atoi(f.Status)
			spd := &types.UserStarDetail{
				StarId:      snowflake.GenID(),
				LikedPostId: f.PostId,
				LikedUserId: f.UserId,
				Status:      int8(status),
				CreatedAt:   utils.TimeStringToGoTime(f.CreatedAt, utils.TimeTemplates),
				UpdatedAt:   utils.TimeStringToGoTime(f.UpdatedAt, utils.TimeTemplates),
			}

			var out model.UserStar
			resultFindUserIsStarPost := db.Model(&model.UserStar{}).Select("star_id").
				Where("liked_post_id = ? and liked_user_id = ?", postId, userId).Find(&out)
			if resultFindUserIsStarPost.RowsAffected < 1 {
				db.Model(&model.UserStar{}).Create(spd)
			} else {
				db.Model(&model.UserStar{}).
					Where("liked_post_id = ? and liked_user_id = ?", postId, userId).
					Updates(types.UserStarDetail{Status: int8(status),
						UpdatedAt: utils.TimeStringToGoTime(f.UpdatedAt, utils.TimeTemplates)})
			}

			c.Del(ctx, getStarRedisKey(postId, userId))
			var likeNum int64
			// 将点赞数量更新进mysql
			db.Model(&model.Post{}).Select("like_num").
				Where("post_id = ?", postId).Find(&likeNum)
			newLikeNum, _ := strconv.ParseInt(c.HGet(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId).Val(), 10, 64)
			if postId != "" {
				db.Model(&model.Post{}).Where("post_id = ?", postId).
					Update("like_num = ?", newLikeNum+likeNum)
			}
			c.HSet(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId, 0)
		}
	}

	zap.L().Info("UpdateStarDetailFromRedisToMySQL 定时任务执行完毕")
	return
}

// 点赞影响post排名，做成推荐文章，但是推荐文章只给200篇，如果后续还有需要，从数据库中

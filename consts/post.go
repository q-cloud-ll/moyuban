package consts

import "errors"

const (
	OrderScore = "score"
)

var (
	PostGetRedisIdsErr = errors.New("查询帖子缓存失败")
	PostListByIdsErr   = errors.New("查询帖子失败")
)

package consts

import "errors"

const (
	OrderScore = "score"
)

var (
	PostGetRedisIdsErr = errors.New("redis.GetPostIdsInOrder(p) return 0 data")
)

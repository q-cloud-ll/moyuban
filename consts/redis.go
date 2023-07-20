package consts

import "errors"

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")

	GetRedisIdsZeroErr = errors.New("redis.GetCommentIdsInOrder(pc) return 0 data")
)

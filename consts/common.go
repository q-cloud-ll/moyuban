package consts

import (
	"errors"

	"golang.org/x/sync/singleflight"
)

// 状态
const (
	StatusOk      = 0 // 正常
	StatusDeleted = 1 // 删除
	StatusReview  = 2 // 待审核
)

const (
	OtherSen      = "./static/dictionary/其他词库.txt"
	Reactionary   = "./static/dictionary/反动词库.txt"
	ViolenceCould = "./static/dictionary/暴恐词库.txt"
	PeoLivelihood = "./static/dictionary/民生词库.txt"
	Porm          = "./static/dictionary/色情词库.txt"
	Corruption    = "./static/dictionary/贪腐词库.txt"
)

// 一些通用sql的错误提示

var (
	InvalidIDErr = errors.New("无效的ID")
)

// jwt
var (
	TokenExpired     = errors.New("token is expired")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
	TokenNotValidYet = errors.New("token not active yet")
	SingleFlight     = &singleflight.Group{}
)

const (
	AccessTokenHeader    = "access_token"
	RefreshTokenHeader   = "refresh_token"
	HeaderForwardedProto = "X-Forwarded-Proto"
	MaxAge               = 3600 * 24
)

// setting
const (
	ConfigEnv         = "GB_CONFIG"
	ConfigDefaultFile = "conf/config.yaml"
	ConfigTestFile    = "conf/config.test.yaml"
	ConfigReleaseFile = "conf/config.release.yaml"
)

// track
const SpanCTX = "span-ctx"

var (
	GetTrackIdErr = errors.New("获取 track id 错误")
)

// upload
const (
	UploadMaxM           = 10
	UploadMaxBytes int64 = 1024 * 1024 * 1024 * UploadMaxM
)

// redis
var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")

	GetRedisIdsZeroErr = errors.New("redis.GetCommentIdsInOrder(pc) return 0 data")
)

// kafka

const (
	KafkaKey     = "disableconsumer"
	TopicLike    = "topic_like"
	TopicComment = "topic_comment"
)

// message
// 消息状态
const (
	StatusUnread   = 0 // 消息未读
	StatusHaveRead = 1 // 消息已读
)

type Type int

// 消息类型
const (
	TypePostComment      Type = iota // 收到帖子评论
	TypePostCommentReply             // 收到他人回复
	TypePostLike                     // 收到点赞
	TypePostFavorite                 // 帖子被收藏
	TypePostRecommend                // 帖子被设为推荐
	TypePostDelete                   // 帖子被删除
)

// 消息主题
const (
	MessageLikePost     = "点赞了你的帖子"
	MessageCommentPost  = "评论了你的帖子"
	MessageReplyComment = "回复了你的评论"
)

// EntityType
const (
	EntityPost    = "post"
	EntityTopic   = "topic"
	EntityComment = "comment"
	EntityUser    = "user"
	EntityCheckIn = "checkIn"
)

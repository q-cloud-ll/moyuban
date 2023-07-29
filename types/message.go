package types

import (
	"project/repository/db/model"
	"time"
)

type MessageResp struct {
	MessageId    int64       `json:"message_id"`
	From         *model.User `json:"from"`    // 消息发送人
	UserId       int64       `json:"user_id"` // 消息接收人编号
	Title        string      `json:"title"`   // 标题
	Content      string      `json:"content"` // 消息内容
	QuoteContent string      `json:"quote_content"`
	Type         int         `json:"type"`
	DetailUrl    string      `json:"detail_url"` // 消息详情url
	ExtraData    string      `json:"extra_data"`
	Status       int         `json:"status"`
	CreateTime   time.Time   `json:"createTime"`
}

type PostLikeExtraData struct {
	PostId     int64 `json:"topic_id"`
	LikeUserId int64 `json:"like_user_id"`
}

type GetMessageReq struct {
	PageInfo
	Type int `json:"type" form:"type"`
}

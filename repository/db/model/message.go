package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	MessageId    int64  `gorm:"not null;index" json:"message_id" form:"message_id"`
	FromId       int64  `gorm:"not null" json:"fromId" form:"fromId"`                              // 消息发送人
	UserId       int64  `gorm:"not null;index:idx_message_user_id;" json:"user_id" form:"user_id"` // 用户编号(消息接收人)
	Type         int    `gorm:"type:int(11);not null" json:"type" form:"type"`                     // 消息类型
	Status       int    `gorm:"type:int(11);not null" json:"status" form:"status"`                 // 状态：0：未读、1：已读
	Title        string `gorm:"size:1024" json:"title" form:"title"`                               // 消息标题
	Content      string `gorm:"type:text;not null" json:"content" form:"content"`                  // 消息内容
	QuoteContent string `gorm:"type:text" json:"quoteContent" form:"quote_content"`                // 引用内容
	ExtraData    string `gorm:"type:text" json:"extraData" form:"extra_data"`                      // 扩展数据
}

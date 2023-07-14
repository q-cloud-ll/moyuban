package model

import "gorm.io/gorm"

// Comment 评论表
type Comment struct {
	gorm.Model
	CommentId int64  `json:"comment_id" gorm:"index;not null;unique;comment:评论id"`
	PostId    int64  `json:"post_id" gorm:"index;not null;comment:帖子id"`
	ReplyId   int64  `json:"reply_id" gorm:"index;not null"`
	Pid       int64  `json:"pid" gorm:"index;not null;comment:父id"`
	UserId    string `json:"user_id" gorm:"index;not null;comment:用户id"`
	Content   string `json:"content" gorm:"not null"`
}

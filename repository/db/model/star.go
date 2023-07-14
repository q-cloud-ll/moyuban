package model

import "gorm.io/gorm"

// UserStar 用户点赞表
type UserStar struct {
	gorm.Model
	StarId      int64  `json:"star_id" gorm:"index;not null;unique; comment:点赞id"`
	LikedUserId string `json:"liked_user_id" gorm:"index;not null; comment:点赞的用户id"`
	LikedPostId string `json:"liked_post_id" gorm:"index;not null; comment:被点赞的帖子id"`
	Status      int8   `json:"status" gorm:"default:1; comment:点赞状态，1点赞，0取消，-1踩"`
}

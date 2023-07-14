package model

import "gorm.io/gorm"

// Follow 关注表
type Follow struct {
	gorm.Model
	FollowId   int64  `json:"follow_id" gorm:"not null; index;comment:关注id"`
	FollowerId string `json:"follower_id" gorm:"not null; index"`
	FolloweeId string `json:"followee_id" gorm:"not null;index"`
}

package model

import "gorm.io/gorm"

// Community 社区分类表
type Community struct {
	gorm.Model
	CommunityId   int64  `json:"community_id" forum:"community_id" gorm:"not null;unique;comment:社区id"`
	CommunityName string `json:"community_name" gorm:"not null;unique;comment:社区名"`
	Introduction  string `json:"introduction,omitempty" gorm:"not null;comment:社区介绍"`
}

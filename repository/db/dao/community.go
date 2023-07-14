package dao

import "gorm.io/gorm"

var _ CommunityModel = (*customCommunityModel)(nil)

type (
	CommunityModel interface {
	}

	customCommunityModel struct {
		*gorm.DB
	}
)

func NewCommunity() CommunityModel {
	return &customCommunityModel{
		DB: NewDBClient(),
	}
}

package dao

import "gorm.io/gorm"

var _ FollowModel = (*customFollowModel)(nil)

type (
	FollowModel interface {
	}

	customFollowModel struct {
		*gorm.DB
	}
)

func NewFollowModel() FollowModel {
	return &customFollowModel{
		DB: NewDBClient(),
	}
}

package dao

import "gorm.io/gorm"

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
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

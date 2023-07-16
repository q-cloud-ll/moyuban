package dao

import "gorm.io/gorm"

var _ UserStarModel = (*customUserStarModel)(nil)

type (
	// UserStarModel is an interface to be customized, add more methods here,
	// and implement the added methods in UserStarModel.
	UserStarModel interface {
	}

	customUserStarModel struct {
		*gorm.DB
	}
)

func NewUserStarModel() UserStarModel {
	return &customUserStarModel{
		DB: NewDBClient(),
	}
}

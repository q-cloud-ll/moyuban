package dao

import "gorm.io/gorm"

var _ UserStarModel = (*customUserStarModel)(nil)

type (
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

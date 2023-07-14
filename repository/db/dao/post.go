package dao

import "gorm.io/gorm"

var _ PostModel = (*customPostModel)(nil)

type (
	PostModel interface {
	}

	customPostModel struct {
		*gorm.DB
	}
)

func NewPostModel() PostModel {
	return &customPostModel{
		DB: NewDBClient(),
	}
}

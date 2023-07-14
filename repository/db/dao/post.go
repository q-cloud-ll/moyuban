package dao

import "gorm.io/gorm"

type (
	PostModel interface {
	}

	customPostModel struct {
		*gorm.DB
	}
)

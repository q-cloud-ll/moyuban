package dao

import "gorm.io/gorm"

var _ CommentModel = (*customUserModel)(nil)

type (
	CommentModel interface {
	}

	customCommentModel struct {
		*gorm.DB
	}
)

func NewCommentModel() CommentModel {
	return &customCommentModel{
		DB: NewDBClient(),
	}
}

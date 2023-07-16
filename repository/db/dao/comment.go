package dao

import "gorm.io/gorm"

var _ CommentModel = (*customUserModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
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

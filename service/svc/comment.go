package svc

import (
	"project/repository/cache"
	"project/repository/db/dao"
)

type CommentServiceContext struct {
	CommentModel dao.CommentModel
	UserModel    dao.UserModel
	CommentCache cache.CommentCache
}

func NewCommentServiceContext() *CommentServiceContext {
	return &CommentServiceContext{
		CommentModel: dao.NewCommentModel(),
		UserModel:    dao.NewUserModel(),
		CommentCache: cache.NewCommentCache(),
	}
}

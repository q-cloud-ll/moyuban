package svc

import (
	"project/repository/cache"
	"project/repository/db/dao"
)

type MessageServiceContext struct {
	MessageModel dao.MessageModel
	MessageCache cache.MessageCache
	UserModel    dao.UserModel
}

func NewMessageServiceContext() *MessageServiceContext {
	return &MessageServiceContext{
		MessageModel: dao.NewMessageModel(),
		MessageCache: cache.NewMessageCache(),
		UserModel:    dao.NewUserModel(),
	}
}

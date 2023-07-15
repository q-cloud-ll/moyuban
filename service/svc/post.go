package svc

import (
	"project/repository/cache"
	"project/repository/db/dao"
)

type PostServiceContext struct {
	PostCache      cache.PostCache
	PostModel      dao.PostModel
	CommunityModel dao.CommunityModel
	UserSvc        *UserServiceContext
}

func NewPostServiceContext() *PostServiceContext {
	return &PostServiceContext{
		PostCache:      cache.NewPostCache(),
		PostModel:      dao.NewPostModel(),
		CommunityModel: dao.NewCommunity(),
		UserSvc:        NewUserServiceContext(),
	}
}

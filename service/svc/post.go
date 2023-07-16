package svc

import (
	"project/repository/cache"
	"project/repository/db/dao"
)

type StarServiceContext struct {
	StarModel dao.UserStarModel
	StarCache cache.StarCache
}

func NewStarServiceContext() *StarServiceContext {
	return &StarServiceContext{
		StarModel: dao.NewUserStarModel(),
		StarCache: cache.NewStarCache(),
	}
}

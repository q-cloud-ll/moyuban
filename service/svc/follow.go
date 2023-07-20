package svc

import (
	"project/repository/cache"
	"project/repository/db/dao"
)

type FollowServiceContext struct {
    FollowModel dao.FollowModel
	  FollowCache cache.FollowCache
}

func NewFollowServiceContext() *FollowServiceContext {
	  return &FollowServiceContext{
		  FollowModel: dao.NewFollowModel(),
		  FollowCache: cache.NewFollowCache(),
	  }
}


package service

import (
    "context"
    "go.uber.org/zap"
    "project/logger"
    "project/service/svc"
)


type FollowSrv struct {
    ctx    context.Context
    svcCtx *svc.FollowServiceContext
    log    *zap.Logger
}

func NewFollowService(ctx context.Context, svcCtx *svc.FollowServiceContext) *FollowSrv {
    return &FollowSrv{
        ctx:    ctx,
        svcCtx: svcCtx,
        log:    logger.Lg,
      }
}

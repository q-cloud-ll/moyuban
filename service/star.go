package service

import (
	"context"
	"go.uber.org/zap"
	"project/logger"
	"project/service/svc"
	"project/types"
	"strconv"
)

type StarSrv struct {
	ctx    context.Context
	svcCtx *svc.StarServiceContext
	log    *zap.Logger
}

func NewStarService(ctx context.Context, svcCtx *svc.StarServiceContext) *StarSrv {
	return &StarSrv{
		ctx:    ctx,
		svcCtx: svcCtx,
		log:    logger.Lg,
	}
}

func (l *StarSrv) StarPostService(uid int64, req *types.StarPostReq) (err error) {
	return l.svcCtx.StarCache.StarPost(l.ctx, strconv.FormatInt(uid, 10), strconv.FormatInt(req.PostId, 10), float64(req.Direction))
}

package service

import (
	"context"
	"project/logger"
	"project/service/svc"
	"project/types"
	"project/utils/msg"
	"strconv"

	"go.uber.org/zap"
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
	err = l.svcCtx.StarCache.StarPost(l.ctx, strconv.FormatInt(uid, 10), req.PostId, float64(req.Direction))
	if err != nil {
		return err
	}
	if req.Direction == 1 {
		pid, _ := strconv.ParseInt(req.PostId, 10, 64)
		msg.SendPostLikeMsgToKafka(l.ctx, pid, uid)
	}
	return
}

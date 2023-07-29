package service

import (
	"context"
	"project/api/render"
	"project/logger"
	"project/service/svc"
	"project/types"
	"project/utils/app"

	"go.uber.org/zap"
)

type MessageSrv struct {
	ctx    context.Context
	svcCtx *svc.MessageServiceContext
	log    *zap.Logger
}

func NewMessageService(ctx context.Context, svcCtx *svc.MessageServiceContext) *MessageSrv {
	return &MessageSrv{
		ctx:    ctx,
		svcCtx: svcCtx,
		log:    logger.Lg,
	}
}

func (l *MessageSrv) GetMessagesSrv(req *types.GetMessageReq) (resp interface{}, total int64, err error) {
	userInfo, err := app.GetUserInfo(l.ctx)
	if err != nil {
		return nil, 0, err
	}
	user, err := l.svcCtx.UserModel.GetUserInfo(l.ctx, userInfo.UID)
	if err != nil {
		return nil, 0, err
	}
	messages, err := l.svcCtx.MessageModel.GetMessageListByPage(l.ctx, req.Type, userInfo.UID, req.Page, req.PageSize)
	if err != nil {
		return nil, 0, err
	}
	_ = l.svcCtx.MessageModel.MarkRead(l.ctx, userInfo.UID)

	resp = render.BuildMessages(user, messages)

	return resp, int64(len(resp.([]types.MessageResp))), err
}

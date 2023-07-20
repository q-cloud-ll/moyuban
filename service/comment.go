package service

import (
	"context"
	"go.uber.org/zap"
	"project/consts"
	"project/logger"
	"project/repository/db/model"
	"project/service/svc"
	"project/types"
	"project/utils/snowflake"
)

type CommentSrv struct {
	ctx    context.Context
	svcCtx *svc.CommentServiceContext
	log    *zap.Logger
}

func NewCommentService(ctx context.Context, svcCtx *svc.CommentServiceContext) *CommentSrv {
	return &CommentSrv{
		ctx:    ctx,
		svcCtx: svcCtx,
		log:    logger.Lg,
	}
}

func (l *CommentSrv) CreateCommentSrv(req *types.PostCommentReq, uid int64) (err error) {
	comment := &model.Comment{
		CommentId: snowflake.GenID(),
		Pid:       req.Pid,
		Content:   req.Content,
		PostId:    req.PostId,
		ReplyId:   req.ReplyId,
		UserId:    uid,
	}
	if req.Pid != 0 {
		err = l.svcCtx.CommentCache.IncrChildrenNum(l.ctx, req.Pid)
		if err != nil {
			zap.L().Info("incr子评论数量失败,err:", zap.Int64("parent comment_id:", comment.CommentId), zap.Error(err))
		}
	}

	err = l.svcCtx.CommentModel.CreateComment(l.ctx, comment)
	if err != nil {
		return
	}

	return l.svcCtx.CommentCache.CreateComment(l.ctx, comment.CommentId, comment.PostId, comment.Pid)
}

// GetPostCommentListSrv 获取帖子下根评论列表
func (l *CommentSrv) GetPostCommentListSrv(req *types.CommentListReq) (resp interface{}, total int64, err error) {
	ids, err := l.svcCtx.CommentCache.GetCommentsIdsInOrder(l.ctx, req)
	if err != nil {
		return nil, 0, err
	}

	if len(ids) == 0 {
		return nil, 0, consts.GetRedisIdsZeroErr
	}

	comments, err := l.svcCtx.CommentModel.GetCommentsListByIds(l.ctx, ids)
	if err != nil {
		return nil, 0, err
	}

	var data []*types.PostCommentResp
	for _, comment := range comments {
		user, err := l.svcCtx.UserModel.GetUserInfo(l.ctx, comment.UserId)
		if err != nil {
			zap.L().Error("GetUserInfo failed", zap.Error(err))
			continue
		}
		childrenNum, _ := l.svcCtx.CommentCache.GetChildrenNum(l.ctx, comment.CommentId)
		star, _ := l.svcCtx.CommentCache.GetCommentStar(l.ctx, comment.CommentId)

		CommentDetail := &types.PostCommentResp{
			CommentId:   comment.CommentId,
			ChildrenNum: childrenNum,
			CreateTime:  comment.CreatedAt,
			Content:     comment.Content,
			LikeNum:     star,
			CommentCreate: &types.CommentCreate{
				CreateById:   user.UserId,
				CreateByName: user.NickName,
				CreateAvatar: user.Avatar,
			},
			CommentReply: &types.CommentReply{},
		}

		data = append(data, CommentDetail)
	}

	total = int64(len(data))

	return
}

// GetChildrenCommentListSrv 获取子评论列表
func (l *CommentSrv) GetChildrenCommentListSrv(req *types.CommentListReq) (resp interface{}, total int64, err error) {
	ids, err := l.svcCtx.CommentCache.GetCommentsIdsInOrder(l.ctx, req)
	if err != nil {
		return nil, 0, err
	}

	if len(ids) == 0 {
		return nil, 0, consts.GetRedisIdsZeroErr
	}

	comments, err := l.svcCtx.CommentModel.GetCommentsListByIds(l.ctx, ids)
	if err != nil {
		return nil, 0, err
	}

	var data []*types.PostCommentResp
	for _, comment := range comments {
		user, err := l.svcCtx.UserModel.GetUserInfo(l.ctx, comment.UserId)
		if err != nil {
			zap.L().Error("GetUserInfo failed", zap.Error(err))
			continue
		}
		replyUser, err := l.svcCtx.CommentModel.GetUserInfoByCommentId(l.ctx, comment.ReplyId)
		if err != nil {
			zap.L().Error("GetUserInfoByCommentId(l.ctx,comment.ReplyId) failed", zap.Error(err))
			continue
		}
		star, _ := l.svcCtx.CommentCache.GetCommentStar(l.ctx, comment.CommentId)

		commentDetail := &types.PostCommentResp{
			CommentId:  comment.CommentId,
			CreateTime: comment.CreatedAt,
			Content:    comment.Content,
			LikeNum:    star,
			CommentCreate: &types.CommentCreate{
				CreateById:   user.UserId,
				CreateByName: user.NickName,
				CreateAvatar: user.Avatar,
			},
			CommentReply: &types.CommentReply{
				ReplyUserId: replyUser.UserId,
				ReplyAvatar: replyUser.Avatar,
				ReplyName:   replyUser.NickName,
			},
		}

		data = append(data, commentDetail)
	}

	total = int64(len(data))

	return
}

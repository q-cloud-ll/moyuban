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
	"strconv"
)

type PostSrv struct {
	ctx    context.Context
	svcCtx *svc.PostServiceContext
	log    *zap.Logger
}

func NewPostService(ctx context.Context, svcCtx *svc.PostServiceContext) *PostSrv {
	return &PostSrv{
		ctx:    ctx,
		svcCtx: svcCtx,
		log:    logger.Lg,
	}
}

// CreatePostSrv 创建帖子服务
func (l *PostSrv) CreatePostSrv(req *types.PostReq) (err error) {
	cid, _ := strconv.ParseInt(req.CommunityId, 10, 64)
	pid := snowflake.GenID()
	post := &model.Post{
		PostId:      pid,
		CommunityId: cid,
		Title:       req.Title,
		Content:     req.Content,
	}

	isExist, err := l.svcCtx.PostModel.JudgeCommunityIsExist(l.ctx, req.CommunityId)
	if isExist {
		return consts.CommunityNotExistErr
	}

	err = l.svcCtx.PostModel.CreatePost(l.ctx, post)
	if err != nil {
		return err
	}
	err = l.svcCtx.PostCache.CreatePost(l.ctx, strconv.FormatInt(pid, 10), req.CommunityId)

	return
}

// GetPostListSrv 获取帖子列表
func (l *PostSrv) GetPostListSrv(req *types.PostListReq) (resp interface{}, total int64, err error) {
	if req.CommunityID == 0 {
		resp, total, err = GetAllPostListSrv(l, req)
	}
	return
}

// GetAllPostListSrv 获取所有帖子列表
func GetAllPostListSrv(l *PostSrv, req *types.PostListReq) (resp interface{}, total int64, err error) {
	ids, err := l.svcCtx.PostCache.GetPostIdsOrder(l.ctx, req)

	if err != nil {
		return nil, 0, err
	}

	if len(ids) == 0 {
		return nil, 0, consts.PostGetRedisIdsErr
	}

	//posts, err := l.svcCtx.PostModel.GetPostListByIds(l.ctx, ids)
	//if err != nil {
	//	return nil, 0, err
	//}
	//
	//l.svcCtx.PostCache
	return
}

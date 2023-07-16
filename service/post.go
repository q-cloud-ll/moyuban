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
	if req.CommunityID == "" {
		resp, total, err = l.GetAllPostListSrv(req)
	} else {
		resp, total, err = l.GetCommunityPostList(req)
	}
	return
}

// GetAllPostListSrv 获取所有帖子列表
func (l *PostSrv) GetAllPostListSrv(req *types.PostListReq) (resp interface{}, total int64, err error) {
	ids, err := l.svcCtx.PostCache.GetPostIdsOrder(l.ctx, req)

	if err != nil {
		return nil, 0, err
	}

	if len(ids) == 0 {
		return nil, 0, consts.PostGetRedisIdsErr
	}

	posts, err := l.svcCtx.PostModel.GetPostListByIds(l.ctx, ids)
	if err != nil {
		return nil, 0, err
	}

	var data []*types.PostDetailResp
	for _, post := range posts.([]*model.Post) {
		user, err := l.svcCtx.UserSvc.UserModel.GetUserInfo(l.ctx, post.AuthorId)
		if err != nil {
			zap.L().Error("GetUserInfo(l.ctx, post.AuthorId) failed",
				zap.Any("author_id", post.AuthorId),
				zap.Error(err))
			continue
		}
		community, err := l.svcCtx.CommunityModel.GetCommunityDetailById(l.ctx, post.CommunityId)
		if err != nil {
			zap.L().Error("GetCommunityDetailById(l.ctx,post.CommunityId) failed",
				zap.Any("author_id", post.AuthorId),
				zap.Error(err))
			continue
		}
		postDetail := &types.PostDetailResp{
			AuthorName: user.NickName,
			Avatar:     user.Avatar,
			//VoteNum:      voteData[idx],
			Post:      post,
			Community: community,
		}
		data = append(data, postDetail)
	}

	total = int64(len(data))

	return data, total, err
}

func (l *PostSrv) GetCommunityPostList(req *types.PostListReq) (resp interface{}, total int64, err error) {
	ids, err := l.svcCtx.PostCache.GetCommunityPostsIdsInorder(l.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if len(ids) == 0 {
		return nil, 0, consts.PostGetRedisIdsErr
	}

	posts, err := l.svcCtx.PostModel.GetPostListByIds(l.ctx, ids)
	if err != nil {
		return nil, 0, consts.PostGetRedisIdsErr
	}

	//voteData, err := FrmGetPostVoteData(ids)
	var data []*types.PostDetailResp
	for _, post := range posts.([]*model.Post) {
		user, err := l.svcCtx.UserSvc.UserModel.GetUserInfo(l.ctx, post.AuthorId)
		if err != nil {
			zap.L().Error("GetUserInfo(l.ctx, post.AuthorId) failed",
				zap.Any("author_id", post.AuthorId),
				zap.Error(err))
			continue
		}
		community, err := l.svcCtx.CommunityModel.GetCommunityDetailById(l.ctx, post.CommunityId)
		if err != nil {
			zap.L().Error("GetCommunityDetailById(l.ctx,post.CommunityId) failed",
				zap.Any("author_id", post.AuthorId),
				zap.Error(err))
			continue
		}

		postDetail := &types.PostDetailResp{
			AuthorName: user.UserName,
			Avatar:     user.Avatar,
			Post:       post,
			Community:  community,
		}
		data = append(data, postDetail)
	}
	total = int64(len(data))
	return
}

func (l *PostSrv) PostDetailSrv(pid string) (resp interface{}, err error) {
	post, err := l.svcCtx.PostModel.GetPostDetailById(l.ctx, pid)
	if err != nil {
		zap.L().Error("GetPostDetailById(pid) failed",
			zap.String("pid", pid),
			zap.Error(err))
		return
	}

	user, err := l.svcCtx.UserSvc.UserModel.GetUserInfo(l.ctx, post.(*model.Post).AuthorId)
	if err != nil {
		zap.L().Error("GetUserInfo(post.(*model.Post).AuthorId) failed",
			zap.String("post.(*model.Post).AuthorId", strconv.FormatInt(post.(*model.Post).AuthorId, 10)),
			zap.Error(err))
		return
	}

	community, err := l.svcCtx.CommunityModel.GetCommunityDetailById(l.ctx, post.(*model.Post).CommunityId)
	if err != nil {
		zap.L().Error("GetCommunityDetailById(l.ctx,post.(*model.Post).CommunityId) failed",
			zap.String("post.(*model.Post).CommunityId", strconv.FormatInt(post.(*model.Post).CommunityId, 10)),
			zap.Error(err))
		return
	}

	resp = &types.PostContentDetailResp{
		Post:      post.(*model.Post),
		User:      user,
		Community: community,
	}

	return
}

//func (l *PostSrv) GetPostVoteData(ids []string) (data []int64, err error) {
//	data = make([]int64, 0, len(ids))
//	for _, id := range ids {
//		v, _ :=
//	}
//}

package service

import (
	"context"
	"project/api/render"
	"project/consts"
	"project/logger"
	"project/repository/db/model"
	"project/service/svc"
	"project/types"
	"project/utils/app"
	"project/utils/snowflake"
	"strconv"
	"strings"

	"github.com/mlogclub/simple/common/jsons"

	"go.uber.org/zap"
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
func (l *PostSrv) CreatePostSrv(req *types.PostReq) (resp interface{}, err error) {
	u, _ := app.GetUserInfo(l.ctx)
	req.Title = strings.TrimSpace(req.Title)
	req.Summary = strings.TrimSpace(req.Summary)
	req.Content = strings.TrimSpace(req.Content)
	cid, _ := strconv.ParseInt(req.CommunityId, 10, 64)
	pid := snowflake.GenID()
	post := &model.Post{
		AuthorId:    u.UID,
		PostId:      pid,
		CommunityId: cid,
		Title:       req.Title,
		Content:     req.Content,
		ContentType: req.ContentType,
		Summary:     req.Summary,
		SourceUrl:   req.SourceUrl,
	}
	if req.Cover != nil {
		post.Cover = jsons.ToJsonStr(req.Cover)
	}
	isExist, err := l.svcCtx.PostModel.JudgeCommunityIsExist(l.ctx, req.CommunityId)
	if isExist {
		return nil, consts.CommunityNotExistErr
	}
	user, _ := l.svcCtx.UserSvc.UserModel.GetUserInfo(l.ctx, u.UID)
	err = l.svcCtx.PostModel.CreatePost(l.ctx, post)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.PostCache.CreatePost(l.ctx, strconv.FormatInt(pid, 10), req.CommunityId)

	return render.BuildPost(post, user, nil), err
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
	for _, post := range posts {
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
	for _, post := range posts {
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
	resp = data
	return
}

func (l *PostSrv) PostDetailSrv(pid string) (resp interface{}, err error) {
	_ = l.svcCtx.PostModel.IncrViewCount(l.ctx, pid)
	post, err := l.svcCtx.PostModel.GetPostDetailById(l.ctx, pid)
	if post == nil || post.Status == consts.StatusDeleted {
		return nil, consts.PostNoFoundErr
	}
	if err != nil {
		zap.L().Error("GetPostDetailById(pid) failed",
			zap.String("pid", pid),
			zap.Error(err))
		return
	}

	user, err := l.svcCtx.UserSvc.UserModel.GetUserInfo(l.ctx, post.AuthorId)
	if err != nil {
		zap.L().Error("GetUserInfo(post.(*model.Post).AuthorId) failed",
			zap.String("post.(*model.Post).AuthorId", strconv.FormatInt(post.AuthorId, 10)),
			zap.Error(err))
		return
	}

	community, err := l.svcCtx.CommunityModel.GetCommunityDetailById(l.ctx, post.CommunityId)
	if err != nil {
		zap.L().Error("GetCommunityDetailById(l.ctx,post.(*model.Post).CommunityId) failed",
			zap.String("post.(*model.Post).CommunityId", strconv.FormatInt(post.CommunityId, 10)),
			zap.Error(err))
		return
	}

	resp = render.BuildPost(post, user, community)

	return
}

func (l *PostSrv) GetEditPostDetail(pid string) (resp interface{}, err error) {
	post, err := l.svcCtx.PostModel.GetPostDetailById(l.ctx, pid)
	if post == nil || post.Status == consts.StatusDeleted {
		return nil, consts.PostNoFoundErr
	}

	return

}

//func (l *PostSrv) GetPostVoteData(ids []string) (data []int64, err error) {
//	data = make([]int64, 0, len(ids))
//	for _, id := range ids {
//		v, _ :=
//	}
//}

package types

import "project/repository/db/model"

type PostReq struct {
	CommunityId string `json:"community_id" validate:"required"`
	Title       string `json:"title"`
	Content     string `json:"content" validate:"required,max=1000"`
}

type PostListReq struct {
	Page        int64  `json:"page" form:"page"`         // 页码
	PageSize    int64  `json:"pageSize" form:"pageSize"` // 每页大小
	CommunityID string `json:"community_id" form:"community_id"`
	Order       string `json:"order" form:"order" example:"score"`
}

type PostDetailResp struct {
	AuthorName       string `json:"author_name"`
	Avatar           string `json:"avatar"`
	VoteNum          int64  `json:"vote_num"`
	*model.Post      `json:"post"`
	*model.Community `json:"community"`
}

// PostContentDetailResp 帖子详情
type PostContentDetailResp struct {
	VoteNum          int64 `json:"vote_num"`
	*model.Post      `json:"post"`
	*model.User      `json:"user"`
	*model.Community `json:"community"`
}

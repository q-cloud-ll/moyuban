package types

import (
	"project/repository/db/model"
	"time"
)

type PostReq struct {
	CommunityId string    `json:"community_id" validate:"required"`
	Title       string    `json:"title" validate:"required,max=20"`
	Content     string    `json:"content" validate:"required,max=1000"`
	ContentType string    `json:"content_type"`
	Summary     string    `json:"summary"`
	SourceUrl   string    `json:"source_url"`
	Cover       *ImageDTO `json:"cover"`
	Tags        []string  `json:"tags"`
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

type PostSimpleRes struct {
	PostId     int64            `json:"post_id"`
	User       *model.User      `json:"user"`
	Community  *model.Community `json:"community"`
	Tags       *[]TagResp       `json:"tags"`
	Title      string           `json:"title"`
	Summary    string           `json:"summary"`
	Cover      *ImageInfo       `json:"cover"`
	SourceUrl  string           `json:"source_url"`
	ViewCount  int64            `json:"view_count"`
	LikeCount  int64            `json:"like_count"`
	CreateTime time.Time        `json:"create_time"`
	Status     int              `json:"status"`
}

type PostResp struct {
	PostSimpleRes
	Content string `json:"content"`
}

type TagResp struct {
	TagId   int64  `json:"tagId"`
	TagName string `json:"tagName"`
}

type ImageInfo struct {
	Url     string `json:"url"`
	Preview string `json:"preview"`
}

type ImageDTO struct {
	Url string `json:"url"`
}

type EditPostDetailReq struct {
	PostId string `form:"post_id" json:"post_id"`
}

type EditPostDetailResp struct {
	PostId  string   `json:"postId"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	cover   string   `json:"cover"`
	Tags    []string `json:"tags"`
}

type SensitiveWord struct {
	Word    string `json:"word"`
	Indexes []int  `json:"indexes"`
	Length  int    `json:"length"`
}

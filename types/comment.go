package types

import "time"

type PostCommentReq struct {
	PostId  int64  `json:"post_id" validate:"required"`
	ReplyId int64  `json:"reply_id"`
	Pid     int64  `json:"pid"`
	Content string `json:"content" validate:"required"`
}

type CommentListReq struct {
	Page     int64  `json:"page" form:"page"`         // 页码
	PageSize int64  `json:"pageSize" form:"pageSize"` // 每页大小
	PostId   string `json:"post_id" form:"post_id"`
	Order    string `json:"order" form:"order" example:"score"`
}

type CdCommentListReq struct {
	Page      int64  `json:"page" form:"page"`         // 页码
	PageSize  int64  `json:"pageSize" form:"pageSize"` // 每页大小
	PostId    string `json:"post_id" form:"post_id"`
	CommentId string `json:"comment_id" form:"comment_id"`
	Order     string `json:"order" form:"order" example:"score"`
}

type StarCommentReq struct {
	Pid       string `json:"pid"`
	PostId    string `json:"post_id"`
	CommentId string `json:"comment_id"`
	Direction int8   `json:"direction"`
}

type PostCommentResp struct {
	CommentId   int64     `json:"comment_id"`
	ChildrenNum int64     `json:"children_num"`
	LikeNum     int64     `json:"like_num"`
	Content     string    `json:"content"`
	CreateTime  time.Time `json:"createTime"`
	*CommentCreate
	*CommentReply
}

type CommentCreate struct {
	CreateById   int64  `json:"create_by_id"`
	CreateByName string `json:"create_by_name"`
	CreateAvatar string `json:"create_avatar"`
}

type CommentReply struct {
	ReplyUserId int64  `json:"reply_user_id"`
	ReplyAvatar string `json:"reply_avatar"`
	ReplyName   string `json:"replyName"`
}

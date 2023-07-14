package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	PostId      int64  `json:"post_id" gorm:"index;not null;unique;comment:帖子id"`
	CommunityId int64  `json:"community_id" gorm:"not null;comment:社区id"`
	CommentId   int64  `json:"comment_id"`
	ReplyId     int64  `json:"reply_id"`
	AuthorId    int64  `json:"author_id" gorm:"not null;comment:作者id"`
	Content     string `json:"content" gorm:"type:longtext;comment:帖子内容"`
	Title       string `json:"title" gorm:"size:500l;comment:帖子标题"`
	Type        int8   `json:"type" gorm:"size:5"`
	LikeNum     int64  `json:"like_num" gorm:"bigint(20)"`
	UnLikeNum   int64  `json:"unLike_num" gorm:"bigint(20)"`
}

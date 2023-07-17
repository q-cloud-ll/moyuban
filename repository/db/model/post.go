package model

import "gorm.io/gorm"

// Post 文章表
type Post struct {
	gorm.Model
	PostId       int64  `json:"post_id" gorm:"index;not null;unique;comment:帖子id"`
	CommunityId  int64  `json:"community_id" gorm:"not null;comment:社区id"`
	Summary      string `gorm:"type:text" json:"summary" form:"summary"`                           // 摘要
	Status       int    `gorm:"type:int(11);index:idx_article_status" json:"status" form:"status"` // 状态
	ViewCount    int64  `gorm:"not null;" json:"viewCount" form:"viewCount"`                       // 查看数量
	CommentCount int64  `gorm:"default:0" json:"commentCount" form:"commentCount"`                 // 评论数量
	ContentType  string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`   // 内容类型：markdown、html
	AuthorId     int64  `json:"author_id" gorm:"not null;comment:作者id"`
	Content      string `json:"content" gorm:"type:longtext;comment:帖子内容"`
	Title        string `json:"title" gorm:"size:500l;comment:帖子标题"`
	Type         int8   `json:"type" gorm:"size:5"`
	LikeNum      int64  `json:"like_num" gorm:"bigint(20)"`
	UnLikeNum    int64  `json:"unLike_num" gorm:"bigint(20)"`
}

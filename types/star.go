package types

import "time"

type StarPostReq struct {
	PostId    int64 `json:"post_id" validate:"required"`
	Direction int8  `json:"direction,string" validate:"required"`
}

type StarPostDetail struct {
	PostId    string `json:"post_id" redis:"post_id"`
	UserId    string `json:"user_id" redis:"user_id"`
	Status    string `json:"status" redis:"status"`
	CreatedAt string `json:"created_at" redis:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" redis:"updated_at"` // 更新时间
}

// UserStarDetail 用户点赞细节 主要有时间问题
type UserStarDetail struct {
	StarId      int64     `json:"star_id"`
	LikedUserId string    `json:"liked_user_id"`
	LikedPostId string    `json:"liked_post_id"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}

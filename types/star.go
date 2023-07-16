package types

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

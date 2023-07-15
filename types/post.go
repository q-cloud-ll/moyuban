package types

type PostReq struct {
	CommunityId string `json:"community_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

type PostListReq struct {
	CommunityID int64  `json:"community_id" form:"community_id"`
	Page        int64  `json:"page" form:"page"`         // 页码
	PageSize    int64  `json:"pageSize" form:"pageSize"` // 每页大小
	Order       string `json:"order" form:"order" example:"score"`
}

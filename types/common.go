package types

type PageInfo struct {
	Page     int64 `json:"page" form:"page"`         // 页码
	PageSize int64 `json:"pageSize" form:"pageSize"` // 每页大小
}

type ActionLink struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

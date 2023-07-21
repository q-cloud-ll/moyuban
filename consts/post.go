package consts

import "errors"

const (
	OrderScore          = "score"
	ContentTypeHtml     = "html"
	ContentTypeMarkdown = "markdown"
	ContentTypeText     = "text"
)

var (
	PostGetRedisIdsErr = errors.New("查询帖子缓存失败")
	PostListByIdsErr   = errors.New("查询帖子失败")
	PostNoFoundErr     = errors.New("文章不存在")
)

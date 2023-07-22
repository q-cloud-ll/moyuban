package consts

import "errors"

const (
	OrderScore          = "score"
	ContentTypeHtml     = "html"
	ContentTypeMarkdown = "markdown"
	ContentTypeText     = "text"
)

var (
	PostGetRedisIdsErr           = errors.New("查询帖子缓存失败")
	PostListByIdsErr             = errors.New("查询帖子失败")
	PostNoFoundErr               = errors.New("文章不存在")
	PostTitleOrContentNoFoundErr = errors.New("文章内容或标题不可为空")
	HasSensitiveWordsErr         = errors.New("内容包含敏感词")
)

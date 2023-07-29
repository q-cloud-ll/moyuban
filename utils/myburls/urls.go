package myburls

import (
	"net/url"
	"project/setting"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

// IsInternalUrl 是否是内部链接
func IsInternalUrl(href string) bool {
	if IsAnchor(href) {
		return true
	}
	u, err := url.Parse(setting.Conf.Host)
	if err != nil {
		zap.L().Error(" url.Parse(setting.Conf.Host)", zap.Error(err))
		return false
	}
	return strings.Contains(href, u.Host)
}

// IsAnchor 是否是锚链接
func IsAnchor(href string) bool {
	return strings.Index(href, "#") == 0
}

func AbsUrl(path string) string {
	return setting.Conf.Host + path
}

func PostUrl(postId int64) string {
	return AbsUrl("/post/" + strconv.FormatInt(postId, 10))
}

func TopicUrl(postId int64) string {
	return AbsUrl("/topic/" + strconv.FormatInt(postId, 10))
}

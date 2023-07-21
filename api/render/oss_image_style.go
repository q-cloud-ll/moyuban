package render

import (
	"project/setting"
	"project/utils/upload"
	"strings"

	"github.com/mlogclub/simple/common/strs"
)

func HandleOssImageStyleAvatar(url string) string {
	if !upload.IsEnabledOss() {
		return url
	}

	return HandleOssImageStyle(url, setting.Conf.QiNiuOssConfig.StyleAvatar)
}

func HandleOssImageStyleDetail(url string) string {
	if !upload.IsEnabledOss() {
		return url
	}
	return HandleOssImageStyle(url, setting.Conf.QiNiuOssConfig.StyleDetail)
}

func HandleOssImageStyleSmall(url string) string {
	if !upload.IsEnabledOss() {
		return url
	}
	return HandleOssImageStyle(url, setting.Conf.QiNiuOssConfig.StyleSmall)
}

func HandleOssImageStylePreview(url string) string {
	if !upload.IsEnabledOss() {
		return url
	}
	return HandleOssImageStyle(url, setting.Conf.QiNiuOssConfig.StylePreview)
}

func HandleOssImageStyle(url, style string) string {
	if strs.IsBlank(style) || strs.IsBlank(url) {
		return url
	}
	if !upload.IsOssImageUrl(url) {
		return url
	}
	if strings.HasSuffix(strings.ToLower(url), ".gif") {
		return url
	}
	sep := setting.Conf.QiNiuOssConfig.StyleSplitter
	if strs.IsBlank(sep) {
		return url
	}
	return strings.Join([]string{url, style}, sep)
}

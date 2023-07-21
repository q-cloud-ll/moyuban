package html

import (
	"project/setting"
	"project/utils/myburls"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mlogclub/simple/common/strs"
	"github.com/mlogclub/simple/common/urls"
)

func xssProtection(htmlContent string) string {
	ugcProtection := bluemonday.UGCPolicy() // 用户生成内容模式
	ugcProtection.AllowAttrs("class").OnElements("code")
	ugcProtection.AllowAttrs("start").OnElements("ol", "ul", "li")
	return ugcProtection.Sanitize(htmlContent)
}

// HandleHtmlContent 处理html内容
func HandleHtmlContent(htmlContent string) string {
	htmlContent = xssProtection(htmlContent)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return htmlContent
	}

	doc.Find("a").Each(func(_ int, selection *goquery.Selection) {
		href := selection.AttrOr("href", "")

		if strs.IsBlank(href) {
			return
		}

		// 不是内部链接
		if !myburls.IsInternalUrl(href) {
			selection.SetAttr("target", "_blank")
			selection.SetAttr("rel", "external nofollow") // 标记站外链接，搜索引擎爬虫不传递权重值

			if setting.Conf.UrlRedirect { // 开启非内部链接跳转
				newHref := urls.ParseUrl(myburls.AbsUrl("/redirect")).AddQuery("url", href).BuildStr()
				selection.SetAttr("href", newHref)
			}
		}

		// 如果a标签没有title，那么设置title
		title := selection.AttrOr("title", "")
		if len(title) == 0 {
			selection.SetAttr("title", selection.Text())
		}
	})

	// 处理图片
	doc.Find("img").Each(func(_ int, selection *goquery.Selection) {
		src := selection.AttrOr("src", "")

		// 处理第三方图片
		if strings.Contains(src, "qpic.cn") {
			src = urls.ParseUrl("/api/img/proxy").AddQuery("url", src).BuildStr()
		}

		// 处理图片样式
		//src = HandleOssImageStyleDetail(src)

		// 处理lazyload
		selection.SetAttr("data-src", src)
		selection.RemoveAttr("src")
	})

	if htmlStr, err := doc.Find("body").Html(); err == nil {
		return htmlStr
	}
	return htmlContent
}

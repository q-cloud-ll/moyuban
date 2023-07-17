package html

import (
	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"strings"
)

// GetHtmlText 获取html文本
func GetHtmlText(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		zap.L().Error("GetHtmlText failed", zap.Error(err))
		return ""
	}
	return doc.Text()
}

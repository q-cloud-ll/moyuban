package render

import (
	"project/consts"
	"project/repository/db/model"
	"project/types"
	"project/utils/html"
	"project/utils/markdown"
)

func BuildPost(post *model.Post, currentUser *model.User, community *model.Community) *types.PostResp {
	if post == nil {
		return nil
	}

	rsp := &types.PostResp{}

	rsp.PostId = post.PostId
	rsp.Title = post.Title
	rsp.Summary = post.Summary
	rsp.SourceUrl = post.SourceUrl
	rsp.ViewCount = post.ViewCount
	rsp.CreateTime = post.CreatedAt
	rsp.Status = post.Status

	rsp.User = currentUser

	rsp.Community = community

	if post.ContentType == consts.ContentTypeMarkdown {
		content := markdown.ToHTML(post.Content)
		rsp.Content = html.HandleHtmlContent(content)
	} else if post.ContentType == consts.ContentTypeHtml {
		rsp.Content = html.HandleHtmlContent(post.Content)
	}

	rsp.Cover = buildImage(post.Cover)
	//if currentUser != nil {
	//	rsp.Favorited = services.FavoriteService.IsFavorited(currentUser.Id, constants.EntityArticle, article.Id)
	//}

	return rsp
}

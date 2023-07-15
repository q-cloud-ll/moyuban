package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/consts"
	"project/service"
	"project/service/svc"
	"project/types"
	"project/utils"
	"project/utils/app"
)

// CreatePostHandler 新建帖子
func CreatePostHandler(c *gin.Context) {
	var req types.PostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("CreatePostHandler param with invalid", zap.Error(err))
		app.ResponseError(c, app.CodeInvalidParam)
		return
	}

	msg, err := utils.Validate(c, &req)
	if err != int64(app.CodeSuccess) {
		zap.L().Error("CreatePostHandler validation error", zap.String("err:", msg))
		app.ResponseErrorWithMsg(c, msg)
		return
	}

	ps := service.NewPostService(c.Request.Context(), svc.NewPostServiceContext())
	if err := ps.CreatePostSrv(&req); err != nil {
		zap.L().Error("CreatePostSrv failed,err:", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, nil)
}

// GetPostListHandler 获取帖子列表接口
func GetPostListHandler(c *gin.Context) {
	p := &types.PostListReq{
		Page:     1,
		PageSize: 10,
		Order:    consts.OrderScore,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler query with invalid", zap.Error(err))
		app.ResponseError(c, app.CodeInvalidParam)
		return
	}

	msg, err := utils.Validate(c, &p)
	if err != int64(app.CodeSuccess) {
		zap.L().Error("GetPostListHandler validation error", zap.String("err:", msg))
		app.ResponseErrorWithMsg(c, msg)
		return
	}

	ps := service.NewPostService(c.Request.Context(), svc.NewPostServiceContext())
	resp, total, errs := ps.GetPostListSrv(p)
	if errs != nil {
		zap.L().Error("GetPostListSrv failed,err:", zap.Error(errs))
		app.ResponseErrorWithMsg(c, errs.Error())
		return
	}

	app.ResponseSuccess(c, app.PageResult{
		List:     resp,
		Total:    total,
		Page:     p.Page,
		PageSize: p.PageSize,
	})
}

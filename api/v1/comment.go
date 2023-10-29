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

func CreateCommentHandler(c *gin.Context) {
	var req types.PostCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("CreateCommentHandler param with invalid", zap.Error(err))
		app.ResponseError(c, app.CodeInvalidParam)
		return
	}
	msg, err := utils.Validate(c, &req)
	if err != int64(app.CodeSuccess) {
		zap.L().Error("CreateCommentHandler validation error", zap.String("err:", msg))
		app.ResponseErrorWithMsg(c, msg)
		return
	}

	user, _ := app.GetUserInfo(c.Request.Context())
	cs := service.NewCommentService(c.Request.Context(), svc.NewCommentServiceContext())
	if err := cs.CreateCommentSrv(&req, user.UID); err != nil {
		zap.L().Error("CreateCommentSrv failed", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, nil)
}

func GetPostCommentListHandler(c *gin.Context) {
	req := &types.CommentListReq{
		Page:     1,
		PageSize: 10,
		Order:    consts.OrderScore,
	}

	if err := c.ShouldBindQuery(req); err != nil {
		zap.L().Error("GetPostCommentListHandler param with invalid", zap.Error(err))
		app.ResponseError(c, app.CodeInvalidParam)
		return
	}

	cs := service.NewCommentService(c.Request.Context(), svc.NewCommentServiceContext())
	resp, total, err := cs.GetPostCommentListSrv(req)
	if err != nil {
		zap.L().Error("GetPostCommentList failed", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, app.PageResult{
		List:     resp,
		PageSize: req.PageSize,
		Page:     req.Page,
		Total:    total,
	})
}

func GetChildrenCommentListHandler(c *gin.Context) {
	req := &types.CommentListReq{
		Page:     1,
		PageSize: 10,
		Order:    consts.OrderScore,
	}

	if err := c.ShouldBindQuery(req); err != nil {
		zap.L().Error("GetChildrenCommentListHandler param with invalid", zap.Error(err))
		app.ResponseError(c, app.CodeInvalidParam)
		return
	}

	cs := service.NewCommentService(c.Request.Context(), svc.NewCommentServiceContext())
	resp, total, err := cs.GetChildrenCommentListSrv(req)
	if err != nil {
		zap.L().Error("GetPostCommentList failed", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, app.PageResult{
		List:     resp,
		PageSize: req.PageSize,
		Page:     req.Page,
		Total:    total,
	})
}

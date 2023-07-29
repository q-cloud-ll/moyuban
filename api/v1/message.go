package v1

import (
	"project/service"
	"project/service/svc"
	"project/types"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetMessagesHandler(c *gin.Context) {
	var req types.GetMessageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		zap.L().Error("GetMessagesHandler param with invalid", zap.Error(err))
		app.ResponseError(c, app.CodeInvalidParam)
		return
	}

	resp, total, err := service.NewMessageService(c.Request.Context(), svc.NewMessageServiceContext()).GetMessagesSrv(&req)
	if err != nil {
		zap.L().Error("GetMessagesSrv failed,err:", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, app.PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     resp,
	})
}

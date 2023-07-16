package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/service"
	"project/service/svc"
	"project/types"
	"project/utils"
	"project/utils/app"
)

func StarPostHandler(c *gin.Context) {
	var req types.StarPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("PostStarHandler param with invalid, err:", zap.Error(err))
		app.ResponseError(c, app.CodeInvalidParam)
		return
	}

	msg, err := utils.Validate(c, &req)
	if err != int64(app.CodeSuccess) {
		zap.L().Error("PostStarHandler validation error", zap.String("err:", msg))
		app.ResponseErrorWithMsg(c, msg)
		return
	}

	u, _ := app.GetUserInfo(c)
	ps := service.NewStarService(c.Request.Context(), svc.NewStarServiceContext())
	if err := ps.StarPostService(u.UID, &req); err != nil {
		zap.L().Error("StarPostService failed", zap.Error(err))
		app.ResponseErrorWithMsg(c, "点赞失败")
		return
	}

	app.ResponseSuccess(c, nil)
}

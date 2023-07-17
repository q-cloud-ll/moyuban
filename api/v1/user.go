package v1

import (
	"project/service"
	"project/service/svc"
	"project/types"
	"project/utils"
	"project/utils/app"
	"project/utils/captcha"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserRegisterHandler(c *gin.Context) {
	var req types.UserRegisterReq
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error("UserRegisterHandler param with invalid", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}
	msg, err := utils.Validate(c, &req)
	if err != int64(app.CodeSuccess) {
		zap.L().Error("UserRegisterHandler validation error", zap.String("err:", msg))
		app.ResponseErrorWithMsg(c, msg)
		return
	}

	us := service.NewUserService(c.Request.Context(), svc.NewUserServiceContext())
	if err := us.UserRegisterSrv(&req); err != nil {
		zap.L().Error("UserRegisterSrv failed,err:", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, app.CodeSuccess)
}

func UserLoginHandler(c *gin.Context) {
	var req types.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("UserLoginHandler param with invalid", zap.Error(err))
		app.ResponseErrorWithMsg(c, app.CodeInvalidParam)
		return
	}

	us := service.NewUserService(c.Request.Context(), svc.NewUserServiceContext())
	resp, err := us.UserLoginSrv(&req)
	if err != nil {
		zap.L().Error("UserLoginHandler failed,err:", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, resp)
}

func UserSendMsgHandler(c *gin.Context) {
	var req types.UserSendMsgReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("UserSendMsgHandler param with invalid", zap.Error(err))
		app.ResponseErrorWithMsg(c, app.CodeInvalidParam)
		return
	}
	msg, code := utils.Validate(c, &req)
	if code != int64(app.CodeSuccess) {
		zap.L().Error("UserSendMsgHandler validation error", zap.String("err:", msg))
		app.ResponseErrorWithMsg(c, msg)
		return
	}
	if err := captcha.SendMessage(c.Request.Context(), req.Mobile, utils.GenerateVerificationCode()); err != nil {
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, app.CodeSuccess)
}

func UserPhoneLoginHandler(c *gin.Context) {
	var req types.UserPhoneLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("UserPhoneLoginHandler param with invalid", zap.Error(err))
		app.ResponseErrorWithMsg(c, app.CodeInvalidParam)
		return
	}
	msg, code := utils.Validate(c, &req)
	if code != int64(app.CodeSuccess) {
		zap.L().Error("UserPhoneLoginHandler validation error", zap.String("err:", msg))
		app.ResponseErrorWithMsg(c, msg)
		return
	}

	us := service.NewUserService(c.Request.Context(), svc.NewUserServiceContext())
	resp, err := us.UserPhoneLoginSrv(&req)
	if err != nil {
		zap.L().Error("UserPhoneLoginSrv failed,err:", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, resp)
}

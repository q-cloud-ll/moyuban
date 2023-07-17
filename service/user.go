package service

import (
	"context"
	"project/consts"
	"project/logger"
	"project/repository/db/model"
	"project/service/svc"
	"project/types"
	"project/utils"
	"project/utils/jwt"
	"project/utils/snowflake"

	"go.uber.org/zap"
)

type UserSrv struct {
	ctx    context.Context
	svcCtx *svc.UserServiceContext
	log    *zap.Logger
}

func NewUserService(ctx context.Context, svcCtx *svc.UserServiceContext) *UserSrv {
	return &UserSrv{
		ctx:    ctx,
		svcCtx: svcCtx,
		log:    logger.Lg,
	}
}

func (l *UserSrv) UserRegisterSrv(req *types.UserRegisterReq) (err error) {
	_, exist, err := l.svcCtx.UserModel.ExistOrNotByUserName(l.ctx, req.UserName)
	if err != nil {
		return err
	}
	if exist {
		err = consts.UserExistErr
	}
	user := &model.User{
		UserName: req.UserName,
		NickName: utils.GenerateRandomNickNameString(),
		UserId:   snowflake.GenID(),
		Status:   model.Active,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		return err
	}

	// 创建用户
	err = l.svcCtx.UserModel.CreateUser(l.ctx, user)
	if err != nil {
		return consts.UserCreateErr
	}

	return
}

func (l *UserSrv) UserLoginSrv(req *types.UserLoginReq) (resp interface{}, err error) {
	var user *model.User
	user, exist, err := l.svcCtx.UserModel.ExistOrNotByUserName(l.ctx, req.UserName)

	if !exist {
		return nil, consts.UserNotExistErr
	}

	if !user.CheckPassword(req.Password) {
		return nil, consts.UserInvalidPasswordErr
	}

	userResp := &types.UserInfoResp{
		UserId:   user.UserId,
		UserName: user.UserName,
		Email:    user.Email,
		NickName: user.NickName,
		Status:   user.Status,
		CreateAt: user.CreatedAt.Unix(),
	}

	return l.TokenNext(user, userResp)
}

func (l *UserSrv) TokenNext(user *model.User, userInfo *types.UserInfoResp) (resp *types.UserTokenData, err error) {
	b := jwt.BaseClaims{
		UID:      user.UserId,
		ID:       user.ID,
		Username: user.UserName,
		Mobile:   user.Mobile,
	}
	accessToken, refreshToken, err := jwt.NewJWT().GenerateToken(b)
	if err != nil {
		return nil, err
	}

	resp = &types.UserTokenData{
		User:         userInfo,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}

func (l *UserSrv) UserPhoneLoginSrv(req *types.UserPhoneLoginReq) (resp interface{}, err error) {
	code, err := l.svcCtx.UserCache.GetPhoneMsg(l.ctx, req.Mobile)
	if err != nil {
		return nil, consts.UserAuthPhoneCodeErr
	}
	if code != req.VerifyCode {
		return nil, consts.UserAuthPhoneCodeErr
	}
	user, exist, err := l.svcCtx.UserModel.ExistOrNotByMobile(l.ctx, req.Mobile)
	var userInfo *types.UserInfoResp
	if !exist {
		userId := snowflake.GenID()
		nickName := utils.GenerateRandomNickNameString()
		status := model.Active

		user := &model.User{
			Mobile:   req.Mobile,
			NickName: nickName,
			UserId:   userId,
			Status:   status,
		}
		if err := l.svcCtx.UserModel.CreateUser(l.ctx, user); err != nil {
			return nil, consts.UserCreateErr
		}
		userInfo = &types.UserInfoResp{
			UserId:   userId,
			UserName: user.UserName,
			Email:    user.Email,
			NickName: nickName,
			Status:   status,
			CreateAt: user.CreatedAt.Unix(),
			Avatar:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
		}
		return l.TokenNext(user, userInfo)
	}
	userInfo = &types.UserInfoResp{
		UserId:   user.UserId,
		UserName: user.UserName,
		Email:    user.Email,
		NickName: user.NickName,
		Status:   user.Status,
		CreateAt: user.CreatedAt.Unix(),
		Avatar:   user.Avatar,
	}

	return l.TokenNext(user, userInfo)
}

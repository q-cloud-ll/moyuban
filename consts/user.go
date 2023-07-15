package consts

import "errors"

var (
	UserExistErr           = errors.New("用户已存在")
	UserNotExistErr        = errors.New("用户不存在")
	UserInvalidPasswordErr = errors.New("用户名或密码错误")
	UserInfoErr            = errors.New("获取用户信息错误")
	UserCreateErr          = errors.New("用户创建失败")
)

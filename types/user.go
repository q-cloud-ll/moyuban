package types

type UserRegisterReq struct {
	UserName   string `json:"user_name" validate:"required,min=6,max=12" label:"用户名"`
	Password   string `json:"password" validate:"required,min=6,max=12" label:"密码"`
	RePassword string `json:"re_password" validate:"required,eqfield=Password,min=6,max=12" label:"重复密码"`
}

type UserLoginReq struct {
	UserName string `json:"user_name" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

type UserPhoneLoginReq struct {
	Mobile     string `json:"mobile" validate:"required,mobile" label:"手机号"`
	VerifyCode string `json:"verify_code" validate:"required" label:"验证码"`
}

type UserSendMsgReq struct {
	Mobile string `json:"mobile" validate:"required,mobile" label:"手机号"`
}

type UserInfoResp struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	UserId   int64  `json:"user_id"`
	CreateAt int64  `json:"create_at"`
	Type     int    `json:"type"`
}

type UserTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

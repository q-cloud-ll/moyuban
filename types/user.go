package types

type UserRegisterReq struct {
	UserName   string `json:"user_name" validate:"required,min=6,max=12" label:"用户名"`
	NickName   string `json:"nick_name" validate:"required,min=6,max=12" label:"昵称"`
	Password   string `json:"password" validate:"required,min=6,max=12" label:"密码"`
	RePassword string `json:"re_password" validate:"required,eqfield=Password,min=6,max=12" label:"重复密码"`
}

type UserLoginReq struct {
	UserName string `json:"user_name" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

type UserInfoResp struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

type UserTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

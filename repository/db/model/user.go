package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// User 用户模型
type User struct {
	gorm.Model
	UserId      int64      `json:"user_id" gorm:"index;unique;not null;comment:用户user_id"`
	LikeNum     int64      `json:"like_num"`
	Birthday    *time.Time `json:"birthday"`
	Gender      int8       `json:"gender"   gorm:"size:1"`
	Type        int8       `json:"type"   gorm:"size:5;comment:是否wx登录"`
	Enable      int        `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	CommentNum  uint32     `json:"comment_num"`
	ArticleNum  uint32     `json:"article_num"`
	Status      string     `json:"status"   gorm:"size:10"`
	Company     string     `json:"company"   gorm:"size:500"`
	WxOpenid    string     `json:"wx_openid"   gorm:"size:500"`
	RealName    string     `json:"real_name" gorm:"size:120"`
	NickName    string     `json:"nick_name" gorm:"size:120"`
	UserName    string     `json:"user_name"  gorm:"size:120"`
	Password    string     `json:"-"  gorm:"size:120"`
	Mobile      string     `json:"mobile"  gorm:"size:120"`
	Email       string     `json:"email" gorm:"size:120"`
	Blog        string     `json:"facebook"   gorm:"size:3000"`
	Avatar      string     `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	Description string     `json:"description"  gorm:"default:Ta很懒，还没有添加简介"`
	Location    string     `json:"location"   gorm:"size:500"`
	School      string     `json:"school"   gorm:"size:500"`
}

const (
	PassWordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

// SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

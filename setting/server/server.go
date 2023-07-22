package server

import (
	"fmt"
	"project/setting"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowServer(router *gin.Engine) {
	address := fmt.Sprintf(":%d", setting.Conf.Port)
	s := initServer(address, router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	zap.L().Info("server run success on ", zap.String("address", address))
	csdn := "https://blog.csdn.net/weixin_51991615"
	fmt.Printf(`
	欢迎使用 go_builder
	当前版本:%s
	简介：基于Gin框架的golang脚手架，内置基本结构，可快速搭建项目
	Up主博客地址：%s
`, setting.Conf.Version, csdn)
	zap.L().Error(s.ListenAndServe().Error())
}

package main

import (
	"fmt"
	"project/logger"
	"project/repository/cache"
	"project/repository/db/dao"
	"project/repository/es"
	"project/repository/track"
	"project/setting"
	"project/utils/snowflake"
	"project/utils/timer"
)

// @title go_builder
// @version 1.0
// @description 基于Go Web 简易脚手架

// @contact.name camellia
// @contact.url https://github.com/q-cloud-ll

// @host 127.0.0.1:8888
// @BasePath /api/v1

func main() {
	//loadingConfig()
	//// 初始化注册路由
	//r := router.SetupRouter()
	//server.RunWindowServer(r)
	//fmt.Println("Starting configuration success...")
	//_ = r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	slice := make([]int, 0, 256)
	for i := 0; i < 256; i++ {
		slice = append(slice, i)
	}
	slice = append(slice, 1)
	fmt.Println(cap(slice))
}

func loadingConfig() {
	setting.Init()
	logger.Init()
	dao.InitMysql()
	cache.InitRedis()
	es.InitEs()
	//rabbitmq.InitRabbitMQ()
	track.InitJaeger()
	snowflake.InitSnowflake()
	timer.TimeTask()
	fmt.Println("Loading configuration success...")
	go scriptStarting()
}

func scriptStarting() {
	// start script

}

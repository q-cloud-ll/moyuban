package main

import (
	"fmt"
	"project/logger"
	"project/repository/cache"
	"project/repository/db/dao"
	"project/repository/es"
	"project/repository/kafka"
	"project/repository/track"
	"project/router"
	"project/setting"
	"project/setting/server"
	"project/utils/snowflake"
	"project/utils/timer"
)

// @title go_builder
// @version 1.0
// @description 基于Gin框架的简易脚手架

// @contact.name camellia
// @contact.url https://github.com/q-cloud-ll

// @host 127.0.0.1:8888
// @BasePath /api/v1

func main() {
	loadingConfig()
	// 初始化注册路由
	r := router.SetupRouter()
	server.RunWindowServer(r)
	fmt.Println("Starting configuration success...")

	_ = r.Run(fmt.Sprintf(":%d", setting.Conf.Port))

}

func loadingConfig() {
	setting.Init()
	logger.Init()
	dao.InitMysql()
	cache.InitRedis()
	es.InitEs()
	kafka.InitKafka()
	//rabbitmq.InitRabbitMQ()
	track.InitJaeger()
	snowflake.InitSnowflake()
	timer.TimeTask()
	fmt.Println("Loading configuration success...")
	go scriptStarting()
	go kafka.StartKafkaConsumer()
}

func scriptStarting() {
	//time.Sleep(2 * time.Second)
	//// start script
	//key := "disableconsumer"
	//topic := "topic1_t"
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//// 使用 Consumer 函数消费指定主题的消息
	//err := kafka.Consumer(ctx, key, topic, func(message *sarama.ConsumerMessage) error {
	//	fmt.Printf("Received message: %s\n", message.Value)
	//	return nil
	//})
	//if err != nil {
	//	fmt.Printf("Failed to consume messages: %s\n", err)
	//}
	//for i := 0; i < 3; i++ {
	//	err := kafka.SendMessage(key, topic, "topic1 send message test"+strconv.Itoa(i))
	//	if err != nil {
	//		fmt.Println("Error sending message", err)
	//		return
	//	}
	//	time.Sleep(1 * time.Second)
	//}
	//
	//ch := make(chan struct{})
	//ch <- struct{}{}
}

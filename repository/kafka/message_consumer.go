package kafka

import (
	"context"
	"encoding/json"
	"project/consts"
	"project/repository/db/dao"
	"project/repository/db/model"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

func StartKafkaConsumer() {
	go ConsumePostLikeGoroutine()

	<-make(chan struct{})
}

func ConsumePostLikeGoroutine() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := Consumer(context.Background(), consts.KafkaKey, consts.TopicLike, func(message *sarama.ConsumerMessage) error {
				var msg model.Message
				_ = json.Unmarshal(message.Value, &msg)
				isExists := dao.NewMessageModel().ExistsMessage(context.Background(), msg.MessageId)
				if isExists {
					return nil
				}
				err := dao.NewMessageModel().CreateMessage(context.Background(), &msg)
				if err == nil {
					zap.L().Info("ConsumePostLikeGoroutine message success", zap.Int64("message_id", msg.MessageId))
				}
				return err
			}); err != nil {
				zap.L().Error("ConsumePostLikeGoroutine failed", zap.Error(err))
			}
		}
	}
}

package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type ConsumerGroupHandler func(message *sarama.ConsumerMessage) error

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := h(msg); err != nil {
			zap.L().Info("消息处理失败",
				zap.String("topic", msg.Topic),
				zap.String("value", string(msg.Value)))
			continue
		}
		sess.MarkMessage(msg, "")
	}

	return nil
}

func MiddlewareConsumerHandler(fn func(message *sarama.ConsumerMessage) error) func(message *sarama.ConsumerMessage) error {
	return func(msg *sarama.ConsumerMessage) error {
		return fn(msg)
	}
}

//	func consumerHandler(message *sarama.ConsumerMessage) error {
//		// 在这里编写处理 Kafka 消息的逻辑
//		message.Value
//		return nil
//	}
func Consumer(ctx context.Context, key, topic string, fn func(message *sarama.ConsumerMessage) error) (err error) {
	kafka, err := GetClient(key)
	if err != nil {
		return
	}

	partitions, err := kafka.Consumer.Partitions(topic)
	if err != nil {
		return
	}

	for _, partition := range partitions {
		// 针对每个分区创建一个对应的分区消费者
		offset, errx := kafka.Client.GetOffset(topic, partition, sarama.OffsetNewest)
		if errx != nil {
			zap.L().Info("获取Offset失败:", zap.Error(errx))
			continue
		}

		pc, errx := kafka.Consumer.ConsumePartition(topic, partition, offset)
		if errx != nil {
			zap.L().Info("获取Offset失败:", zap.Error(errx))
			return
		}
		// 从每个分区都消费消息
		go func(consumer sarama.PartitionConsumer) {
			defer func() {
				if err := recover(); err != nil {
					zap.L().Error("消费kafka信息发生panic,err:%s", zap.Any("err:", err))
				}
			}()

			defer func() {
				err := pc.Close()
				if err != nil {
					zap.L().Error("消费kafka信息发生panic,err:%s", zap.Any("err:", err))
				}
			}()

			for {
				select {
				case msg := <-pc.Messages():
					fmt.Println(msg.Value, "yes")
					err := MiddlewareConsumerHandler(fn)(msg)
					if err != nil {
						return
					}
				case <-ctx.Done():
					return
				}
			}

		}(pc)
	}
	return nil
}

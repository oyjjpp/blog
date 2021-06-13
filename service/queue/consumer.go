package queue

import (
	"context"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	ctx context.Context
}

// 校验接口是否实现
var _ sarama.ConsumerGroupHandler = &Consumer{}

func (consumer *Consumer) Setup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) Cleanup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// for msg := range claim.Messages() {
	// 	message := string(msg.Value)
	// 	consumer.handle(message)
	// 	session.MarkMessage(msg, "")
	// }
	// return nil

	for {
		select {
		case data, ok := <-claim.Messages():
			if !ok {
				log.Printf("ConsumeClaim fail:%v \n", data)
				time.Sleep(10 * time.Second)
				continue
			}
			message := string(data.Value)
			consumer.handle(message)
			session.MarkMessage(data, "")
		case <-consumer.ctx.Done():
			log.Printf("ConsumeClaim consume_message exit")
			return nil
		}
	}
}

func (consumer *Consumer) handle(message string) {
	log.Printf("consumer message:%v \n", message)
}

// kafka消费者初始化
func ConsumerInit(ctx context.Context) {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokerList[:], "consumer_topic-study", config)
	if err != nil {
		log.Printf("ConsumerInit fail : %v\n", err)
		panic(err)
	}

	consumer := Consumer{ctx}
	go func() {
		for {
			err := client.Consume(ctx, []string{"topic-study"}, &consumer)
			if err != nil {
				log.Printf("ConsumerInit Consume fail:%v \n", err)
				time.Sleep(time.Second * 5)
			}
		}
	}()
}

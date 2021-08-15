package queue

import (
	"context"
	"fmt"
	"time"

	cluster "github.com/bsm/sarama-cluster"
)

func RunConsumer(ctx context.Context) {
	//ctx := context.Background()
	consumer := getKafkaConsumer(brokerList[:], []string{"topic-study"}, "topic-study_consumer")
	for i := 0; i < 10; i++ {
		go Run(ctx, consumer)
	}
}

func Run(ctx context.Context, consumer *cluster.Consumer) {
	for {
		select {
		case data, ok := <-consumer.Messages():
			if !ok {
				time.Sleep(10 * time.Second)
				continue
			}
			//logger.Ix(ctx, "consume_message", "message:%s", string(data.Value))
			//process(ctx, messageService, data.Value)

			consumer.MarkOffset(data, "")
		case <-ctx.Done():
			fmt.Println("consume_message exit")
			return
		}
	}
}

func getKafkaConsumer(brokerList, topicList []string, groupName string) *cluster.Consumer {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	consumer, err := cluster.NewConsumer(brokerList, groupName, topicList, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for err := range consumer.Errors() {
			fmt.Printf("get error:%v\n", err)
		}
	}()

	go func() {
		for ntf := range consumer.Notifications() {
			fmt.Printf("get rebalanced notification:%v\n", ntf)
		}
	}()

	return consumer
}

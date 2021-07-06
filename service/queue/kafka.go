package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

type syncProducerPool struct {
	ctx    context.Context
	client sarama.SyncProducer
}

var defaultProducer *syncProducerPool
var brokerList = [...]string{"192.168.124.28:9092", "192.168.124.28:9093", "192.168.124.28:9094"}

// kafka初始化
func ProducerInit(ctx context.Context) {

	log.Println("producer")

	defaultProducer = new(syncProducerPool)
	defaultProducer.ctx = ctx

	config := sarama.NewConfig()
	//config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Version = sarama.V0_10_2_0

	config.ClientID = "1" // 设置客户端ID

	//config.Producer.Retry =  // 重试次数

	// 连接kafka
	//logger.I("syncProducerPool.NewWorker.start", "Consumer:%v;", mkConst.WECHAT_EVENT_TOPIC)
	client, err := sarama.NewSyncProducer(brokerList[:], config)
	if err != nil {
		panic(err)
	}

	log.Println(client)

	defaultProducer.client = client
	return
}

// kafka发送消息
func SendMessage(topic string, message interface{}) error {
	msg := sarama.ProducerMessage{}
	msg.Topic = topic
	strEi, err := json.Marshal(message)
	if err != nil {
		return err
	}
	msg.Value = sarama.StringEncoder(strEi)
	// 日志时间
	msg.Timestamp = time.Now()

	log.Println(defaultProducer.client)

	log.Printf("syncProducerPool.SendMessage begin to send message:%v \n", msg)

	partition, offset, err := defaultProducer.client.SendMessage(&msg)
	log.Printf("partition:%v offset:%d \n", partition, offset)
	if err != nil {
		log.Printf("syncProducerPool.SendMessage.fail msg;%v, error:%v", msg, err)
		return err
	}

	// 定制分区
	return nil
}

// 关闭kafka连接
func CloseKafka() {
	err := defaultProducer.client.Close()
	log.Println(err)
}

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

func Test() {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
}

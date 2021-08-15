package queue

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type syncProducerPool struct {
	ctx    context.Context
	client sarama.SyncProducer
}

var defaultProducer *syncProducerPool
var brokerList = [...]string{"192.168.124.60:9002", "192.168.124.60:9003", "192.168.124.60:9004"}

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

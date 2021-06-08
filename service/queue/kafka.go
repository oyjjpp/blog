package queue

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

type syncProducerPool struct {
	ctx    context.Context
	client sarama.SyncProducer
}

var defaultProducer *syncProducerPool

func ProducerInit(ctx context.Context) {

	log.Println("producer")

	defaultProducer = new(syncProducerPool)
	defaultProducer.ctx = ctx

	config := sarama.NewConfig()
	//config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Version = sarama.V0_10_2_0
	//config.Producer.

	// 连接kafka
	//logger.I("syncProducerPool.NewWorker.start", "Consumer:%v;", mkConst.WECHAT_EVENT_TOPIC)
	client, err := sarama.NewSyncProducer([]string{"192.168.124.28:9092", "192.168.124.28:9093", "192.168.124.28:9094"}, config)
	if err != nil {
		panic(err)
	}

	log.Println(client)

	defaultProducer.client = client
	return
}

func SendMessage(topic string, message interface{}) error {
	msg := sarama.ProducerMessage{}
	msg.Topic = topic
	strEi, err := json.Marshal(message)
	if err != nil {
		return err
	}
	msg.Value = sarama.StringEncoder(strEi)

	log.Println(defaultProducer.client)

	log.Printf("syncProducerPool.SendMessage begin to send message:%v", msg)

	partition, offset, err := defaultProducer.client.SendMessage(&msg)
	log.Printf("partition:%v offset:%d", partition, offset)
	if err != nil {
		log.Printf("syncProducerPool.SendMessage.fail msg;%v, error:%v", msg, err)
		return err
	}

	return nil
}

func CloseKafka() {
	err := defaultProducer.client.Close()
	log.Println(err)
}

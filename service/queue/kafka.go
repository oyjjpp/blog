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
	defaultProducer = new(syncProducerPool)
	defaultProducer.ctx = ctx

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 连接kafka
	//logger.I("syncProducerPool.NewWorker.start", "Consumer:%v;", mkConst.WECHAT_EVENT_TOPIC)
	client, err := sarama.NewSyncProducer([]string{"192.168.124.28:9092"}, config)
	if err != nil {
		panic(err)
	}

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
	log.Printf("syncProducerPool.SendMessage begin to send message:%v", msg)
	_, _, err = defaultProducer.client.SendMessage(&msg)
	if err != nil {
		log.Printf("syncProducerPool.SendMessage.fail msg;%v, error:%v", msg, err)
		return err
	}

	return nil
}

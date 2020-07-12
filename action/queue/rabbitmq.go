package queue

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func Rabbitmap(ctx *gin.Context) {
}

func connection() *amqp.Connection {
	// 连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	return conn
}

func sample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sampleConsumption(&wg)

	conn := connection()
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// 创建交换器
	if err := channel.ExchangeDeclare(
		"e1",
		"direct",
		true,
		false,
		false,
		true,
		nil); err != nil {
		panic(err)
	}

	// 创建路由器
	if _, err := channel.QueueDeclare(
		"q1",
		true,
		false,
		false,
		true,
		nil); err != nil {
		panic(err)
	}

	// 绑定队列
	if err := channel.QueueBind("q1", "q1Key", "e1", true, nil); err != nil {
		panic(err)
	}

	// mandatory true 未找到队列返回给消费者
	returnChan := make(chan amqp.Return, 0)
	channel.NotifyReturn(returnChan)

	// pushlish
	if err := channel.Publish(
		"e1",
		"q1Key",
		true,
		false,
		amqp.Publishing{
			Timestamp:   time.Now(),
			ContentType: "text/plain",
			Body:        []byte("Hello Golang and AMQP(Rabbitmq)!"),
		}); err != nil {
		panic(err)
	}

	for v := range returnChan {
		fmt.Printf("Return %#v\n", v)
	}

	wg.Wait()
}

func sampleConsumption(wg *sync.WaitGroup) {
	// 创建链接
	conn := connection()
	defer conn.Close()

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// 消费信息
	deliveries, err := channel.Consume("q1", "any", false, false, false, true, nil)
	if err != nil {
		panic(err)
	}
	// 取一条消息
	if v, ok := <-deliveries; ok {
		if err := v.Ack(true); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("channel close")
	}
	wg.Done()
}

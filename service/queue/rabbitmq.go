// 消息队列
// rabbitmp
// AMQP(高级消息队列协议)是一个进程间传递异步消息的网络协议
// 使用Erlang语言实现
// 架构的主要角色:生产者、消费者、交换器、队列

/*
交换器：消息代理待服务中勇于把消息路由到队列的组件
队列：用来存储消息的数据结构，位于硬盘或内存中
绑定：一套规则，勇于告诉交换器消息应该被存储到哪个队列
*/
package queue

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

/*
绑定：
Rabbitmq中需要“路由键”和“绑定键”联合使用才能使生产者成功投递到队列中
RoutingKey：生产者发送给交换器绑定的key
BindingKey：交换器和队列绑定的Key

交换器（rabbitmq共有四种交换器）：
fanout：把消息投递到所有与此交换器绑定的队列中
direct：把消息投递到BindingKey和RoutingKey完全匹配的队列中
topic：BindingKey中存在两种特殊字符
*：匹配零个或多个单词
#：匹配一个单词
header：不依赖与RoutingKey而是通过消息体中的headers属性来进行匹配绑定
通过headers中的key和BindingKey完全匹配
*/

/*
Golang中创建rabbitmq生产者基本步骤
1、连接Connection
2、创建Channel
3、创建或连接一个交换器
4、创建活连接一个队列
5、交换器绑定队列
6、投递消息
7、关闭Channel
8、关闭Connection
*/

type Rabbitmq struct {
}

// connect
// 连接
func (r *Rabbitmq) connect() *amqp.Channel {
	// TODO
	// 1、通过什么协议进行连接
	// 2、如何保持长连接
	// 3、是否需要连接池
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// 选择管道
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return channel
}

// createExchange
// 创建一个交换器
func (r *Rabbitmq) createExchange(channel *amqp.Channel) {
	// name:交换器名称
	// kind:交换器类型
	// durable:持久化
	// autoDelete : 是否自动删除
	// internal :是否是内置交换器
	// noWait: 是否等待服务器确认
	// args：其他配置
	if err := channel.ExchangeDeclare(
		"e1", "direct", true, false, false, true, nil); err != nil {
		panic(err)
	}
	// channel.ExchangeDeclarePassive
	// channel.ExchangeBind
	// channel.ExchangeDelete
	// channel.ExchangeUnbind
}

// createQueue
// 创建一个队列
func (r *Rabbitmq) createQueue(channel *amqp.Channel) {
	// name:队列名称
	// durable:持久化
	// autoDelete : 是否自动删除
	// exclusive : 排他
	// noWait: 是否等待服务器确认
	// args：其他配置
	_, err := channel.QueueDeclare("q1", true, false, false, true, nil)
	if err != nil {
		panic(err)
	}
	//channel.QueueDeclarePassive
	// channel.QueueBind
	// channel.QueueUnbind
	// channel.QueueDelete
	// channel.QueueInspect
	// channel.QueuePurge
}

// bindQueue
// 绑定交换器和队列
func (r *Rabbitmq) bindQueue(channel *amqp.Channel) {
	// name:队列名称
	// key BindingKey 根据交换器类型来设定【direct】
	// exchange 交换器名称
	// noWait: 是否等待服务器确认
	// args：其他配置
	err := channel.QueueBind("q1", "q1Key", "e1", true, nil)
	if err != nil {
		panic(err)
	}
}

// bindExchange
// 绑定交换器
func (r *Rabbitmq) bindExchange(channel *amqp.Channel) {
	// destination:目的交换器
	// key RoutingKey 路由键【生产者与交换器】
	// source 源交换器
	// noWait: 是否等待服务器确认
	// args：其他配置
	err := channel.ExchangeBind("dest", "q1Key", "src", false, nil)
	if err != nil {
		panic(err)
	}
}

func (r *Rabbitmq) publish(channel *amqp.Channel) {
	// exchange : 交换器名称
	// key RouterKey
	// mandatory : 是否为无法路由的消息进行返回处理
	// immediate : 是否对路由到无消费者队列的消息进行返回处理(Rabbitmq3.0废弃)
	// msg ： 消息体
	err := channel.Publish("e1", "q1Key", true, false,
		amqp.Publishing{
			Timestamp:   time.Now(),
			ContentType: "text/plain",
			Body:        []byte("Hello Golang and Amqp(Rabbitmp)"),
		})
	if err != nil {
		panic(err)
	}
}

/*
Rabbitmq 消费方式 推模式和拉模式
推模式：通过持续订阅的方式来消费消息
推模式是通过持续订阅的方式来消费消息
拉模式：相对来说简单，是有消费者主动拉去消息来消费
*/

// consume
// 推模式消费信息
func (r *Rabbitmq) consume(channel *amqp.Channel) {
	// queue string 队列
	// consumer string 消费者则名称
	// autoAck bool 是否确认消费
	// exclusive bool, 排它
	// noLocal bool,
	// noWait bool,
	// args Table)
	// (<-chan Delivery, error)
	deliveries, err := channel.Consume(
		"q1",
		"any",
		false,
		false,
		false,
		true,
		nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(deliveries)
	// act 设置false需要手动进行act消费
	if v, ok := <-deliveries; ok {
		if err := v.Ack(true); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("channel Close")
	}
}

// consumeGet
// 拉模式消费信息
func (r *Rabbitmq) consumeGet(channel *amqp.Channel) {
	channel.Get("q1", false)
}

/*
解耦应用

*/

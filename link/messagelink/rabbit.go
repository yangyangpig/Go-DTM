package messagelink

import (
	"github.com/streadway/amqp"
	"github.com/uber/jaeger-client-go/crossdock/log"
	"sync"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

var (
	once sync.Once
	obj  *RabbitMQ
)

//单例模式，减少结构体声明
func NewRabbitMQFactory(c *amqp.Connection) *RabbitMQ {
	once.Do(func() {
		obj = new(RabbitMQ)
		obj.conn = c
	})
	return obj
}
func FailError(err error, msg string) {
	if err != nil {
		log.Printf("err%s", msg)
	}
}

func (this *RabbitMQ) ChoseChannel() *RabbitMQ {
	ch, err := this.conn.Channel()
	this.ch = ch
	FailError(err, "get channel failed")
	return this
}

func (this *RabbitMQ) CreateExchange(name string, kind string, durable, autoDelete, internal, noWait bool, arg map[string]interface{}) error {
	err := this.ch.ExchangeDeclare(
		name,
		kind,
		durable,
		autoDelete,
		internal,
		noWait,
		arg,
	)
	FailError(err, "get channel failed")
	return err
}

func (this *RabbitMQ) Publish(exchange, key string, mandatory, immediate bool, contentType string, body []byte) {
	err := this.ch.Publish(
		exchange,
		key, //route_key
		mandatory,
		immediate,
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, //设置消息持久化
			ContentType: contentType,
			Body:        body,
		})
	FailError(err, "Failed to publish a message")
}

func (this *RabbitMQ) CreateQueue(queueName string, durable, autoDelete, exclusive, noWait bool, args map[string]interface{}) (*RabbitMQ, error) {
	q, err := this.ch.QueueDeclare(
		queueName,
		durable,
		autoDelete,
		exclusive,
		noWait,
		args,
	)
	FailError(err, "Failed to create queue ")
	this.q = q
	return this, nil
}

func (this *RabbitMQ) CreateConsume(customName string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{}) (<-chan amqp.Delivery, error) {
	msgs, err := this.ch.Consume(
		this.q.Name,
		customName,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)
	FailError(err, "Failed to register a consumer")
	return msgs, err
}

func (this *RabbitMQ) ToBind(routeKeys []string, exchange string, noWait bool, args map[string]interface{}) {
	for _, v := range routeKeys {
		err := this.ch.QueueBind(
			this.q.Name, // queue name
			v,           // binding key
			"exchange",  // exchange
			noWait,
			args)
		FailError(err, "Failed to bind a queue")
	}
}

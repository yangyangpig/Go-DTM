package rabbitmq_dirver

//每次引入rabbit驱动时候，就引入这个文件即可,因为这个文件已经把rabbit的driver注册到sql的driver里面，而这个driver已经实现了
//接口sql里面的driver，当sql调用open时候，会调用这个文件的open方法，实现具体的连接。

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQDriver struct{}

func Open(drivename string, dsn string) (*channelPool, error) {
	//dsn格式如下"amqp://admin:123456@47.106.120.121:5672/admin"
	factory := func() (*amqp.Connection, error) { return amqp.Dial(dsn) }
	conn, err := NewChannelPool(3, 100, factory)

	if err != nil {
		log.Fatal("connect rabbitmq failed", err)
		return nil, err
	}

	return conn, nil
}

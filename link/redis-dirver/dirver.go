package redis_dirver

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisDirver struct {
}

func Open(driverName string, dsn string) (redis.Conn, error) {
	//dns :host:port
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", dsn)
		if err != nil {
			return nil, err
		}
		//TODO 阿里云需要做鉴权
		//if len(password) > 0 {
		//  c.Do("AUTH", password)
		//}
		return
	}

	// 初始化连接池
	p := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}

	c := p.Get()
	defer c.Close()
	return c, nil
}

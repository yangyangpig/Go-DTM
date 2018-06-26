package cachelink

import (
	"Go-DMT/Go-DTM/link"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"sync"
)

type RedisCli struct {
	Conn redis.Conn
}

var (
	once sync.Once
	obj  *RedisCli
)

func NewRedisFactory() *RedisCli {
	once.Do(func() {
		obj = new(RedisCli)
		obj.Conn = link.Alias.RD //由于这里redis用了redisgo的池子了，不需要用自己的池子
	})

	return obj
}

func (this *RedisCli) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}
	//conn := this.Conn.Get()

	defer this.Conn.Close()
	return this.Conn.Do(commandName, args...)
}

func (this *RedisCli) LRange(key string, beginID int, endID int) ([]string, error) {
	ret, err := redis.Strings(this.Do("LRANGE", key, beginID, endID))
	return ret, err
}

func (this *RedisCli) RPOP(key string) (string, error) {
	ret, err := redis.String(this.Do("RPOP", key))
	return ret, err
}

func (this *RedisCli) LLEN(key string) (int64, error) {
	ret, err := redis.Int64(this.Do("LLEN", key))
	return ret, err
}

func (this *RedisCli) LTRIM(key string, beginID int, endID int) (string, error) {
	ret, err := redis.String(this.Do("LTRIM", key, beginID, endID))
	return ret, err
}
func (this *RedisCli) LPush(key string, value string) (int, error) {
	ret, err := redis.Int(this.Do("LPUSH", key, value))
	return ret, err
}
func (this *RedisCli) RPush(key string, value string) (int, error) {
	ret, err := redis.Int(this.Do("RPUSH", key, value))
	return ret, err
}

func (this *RedisCli) LPOP(key string) (string, error) {
	ret, err := redis.String(this.Do("LPOP", key))
	return ret, err
}

func (this *RedisCli) HSET(key string, field string, value string) (int64, error) {
	ret, err := redis.Int64(this.Do("HSET", key, field, value))
	return ret, err
}

func (this *RedisCli) HMSETSTR(args ...interface{}) (string, error) {
	ret, err := redis.String(this.Do("HMSET", args...))
	return ret, err
}

func (this *RedisCli) HDEL(args ...interface{}) (int64, error) {
	ret, err := redis.Int64(this.Do("HDEL", args...))
	return ret, err
}

func (this *RedisCli) HGET(key string, field string) (string, error) {
	ret, err := redis.String(this.Do("HGET", key, field))
	return ret, err
}

func (this *RedisCli) HGETInt64(key string, field string) (int64, error) {
	ret, err := redis.Int64(this.Do("HGET", key, field))
	return ret, err
}

func (this *RedisCli) HLEN(key string) (int64, error) {
	ret, err := redis.Int64(this.Do("HLEN", key))
	return ret, err
}

// Ping
func (this *RedisCli) PING() (bool, error) {
	ret, err := redis.String(this.Do("PING"))
	if err != nil {
		return false, err
	}
	if ret == "PONG" {
		return true, nil
	}
	return true, nil
}

// 开始事务
func (this *RedisCli) MULTI() (string, error) {
	ret, err := redis.String(this.Do("MULTI"))
	return ret, err
}

// 提交事务
func (this *RedisCli) EXEC() ([]string, error) {
	ret, err := redis.Strings(this.Do("EXEC"))
	return ret, err
}

// 取消事务
func (this *RedisCli) DISCARD() (string, error) {
	ret, err := redis.String(this.Do("DISCARD"))
	return ret, err
}

func (this *RedisCli) Del(args ...interface{}) (int64, error) {
	ret, err := redis.Int64(this.Do("Del", args...))
	return ret, err
}

func (this *RedisCli) HGETALL(key string) ([]string, error) {
	ret, err := redis.Strings(this.Do("HGETALL", key))
	return ret, err
}

func (this *RedisCli) HGETALLInt64(key string) (map[string]int64, error) {
	ret, err := redis.Int64Map(this.Do("HGETALL", key))
	return ret, err
}

func (this *RedisCli) HGETALLString(key string) (map[string]string, error) {
	ret, err := redis.StringMap(this.Do("HGETALL", key))
	return ret, err
}

func (this *RedisCli) GEOADD(key string, latitude float32, longitude float32, name int64) (int, error) {
	ret, err := redis.Int(this.Do("GEOADD", key, latitude, longitude, name))
	return ret, err
}

func (this *RedisCli) GEODIST(key string, name1 int64, name2 int64) (int64, error) {
	ret, err := redis.Int64(this.Do("GEODIST", key, name1, name2, " m"))
	return ret, err
}

func (this *RedisCli) GEORADIUSBYMEMBER(key string, name int64, distance int) ([]string, error) {
	ret, err := redis.Strings(this.Do("GEORADIUSBYMEMBER", key, name, distance, "km"))
	return ret, err
}

func (this *RedisCli) GEODEL(key string, name int64) (int, error) {
	ret, err := redis.Int(this.Do("ZREM", key, name))
	return ret, err
}

func (this *RedisCli) INCR(key string) (int64, error) {
	ret, err := redis.Int64(this.Do("INCR", key))
	return ret, err
}

func (this *RedisCli) HINCR(key string, field string) (int64, error) {
	ret, err := redis.Int64(this.Do("HINCRBY", key, field, 1))
	return ret, err
}

func (this *RedisCli) GET(key string) (string, error) {
	ret, err := redis.String(this.Do("GET", key))
	return ret, err
}

func (this *RedisCli) EXPIRE(key string, second int64) (int, error) {
	ret, err := redis.Int(this.Do("EXPIRE", key, second))
	return ret, err
}

func (this *RedisCli) SADD(args ...interface{}) (int64, error) {
	ret, err := redis.Int64(this.Do("SADD", args...))
	return ret, err
}
func (this *RedisCli) SREM(args ...interface{}) (int, error) {
	ret, err := redis.Int(this.Do("SREM", args...))
	return ret, err
}

func (this *RedisCli) SMEMBER(key string) ([]int64, error) {
	ret, err := redis.Int64s(this.Do("SMEMBERS", key))
	return ret, err
}

func (this *RedisCli) SISMEMBER(key string, member int64) (int, error) {
	ret, err := redis.Int(this.Do("SISMEMBER", key, member))
	return ret, err
}

func (this *RedisCli) ZADD(args ...interface{}) (int64, error) {
	ret, err := redis.Int64(this.Do("ZADD", args...))
	return ret, err
}

func (this *RedisCli) ZRANGE(args ...interface{}) ([]int64, error) {
	ret, err := redis.Int64s(this.Do("ZRANGE", args...))
	return ret, err
}

// ZREVRANGE mid_test 0 -1 WITHSCORES
func (this *RedisCli) ZREVRANGE(args ...interface{}) ([]int64, error) {
	ret, err := redis.Int64s(this.Do("ZREVRANGE", args...))
	return ret, err
}

func (this *RedisCli) SET(key string, value interface{}) error {
	ret, err := redis.String(this.Do("SET", key, value))

	if err != nil {
		return err
	}
	if ret != "OK" {
		return fmt.Errorf("redis set ret err: %s", ret)
	}

	return err
}

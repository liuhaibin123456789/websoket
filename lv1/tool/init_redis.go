package tool

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// RDB redis连接
var RDB redis.Conn

func InitRedis() error {
	conn, err := redis.Dial("tcp", "localhost:6379", redis.DialWriteTimeout(time.Second), redis.DialReadTimeout(time.Second))
	if err != nil {
		return err
	}
	RDB = conn

	return nil
}

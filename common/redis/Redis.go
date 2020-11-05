/**
 * @Author: xianxiong
 * @Date: 2020/11/4 9:44
 */

package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	pool *redis.Pool
)

//redis连接池
func InitRedis() {
	network := viper.GetString("redis.network")
	address := viper.GetString("redis.address")
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0, //0表示没有限制
		IdleTimeout: 1 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(network, address)
		},
	}
}

//string 操作
func Set(key string, value interface{}) bool {
	// 获取连接
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", key, value)
	if err != nil {
		log.Println("set a value error,", err)
		return false
	}
	return true
}

func Get(key string) string {
	conn := pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("get", key))
	if err != nil {
		log.Println("get a value error,", err)
		return ""
	}

	return value
}

func Del(key string) bool {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("del", key)
	if err != nil {
		log.Println("del a key error, ", err)
		return false
	}
	return true
}

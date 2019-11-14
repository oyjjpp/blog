package service

import (
	"errors"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisPool struct {
	pool     []*redis.Pool
	pool_num int
}

//初始化redis连接池
func (rp *RedisPool) InitRedisPool(ipPort []string) {
	rp.pool = make([]*redis.Pool, 0, 30)

	for i := 0; i < len(ipPort); i++ {
		//拆分链接配置
		//127.0.0.1:6379
		ipInfo := strings.Split(ipPort[i], "@")
		rPool := &redis.Pool{
			//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
			MaxIdle: 1024,
			//
			MaxActive:   2048,
			IdleTimeout: 180 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", ipInfo[0])
				if err != nil {
					return nil, err
				}
				//如果配置了登录校验
				if len(ipInfo) == 2 {
					if _, err := c.Do("AUTH", ipInfo[1]); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, err
			},
		}
		rp.pool = append(rp.pool, rPool)
	}
	rp.pool_num = len(ipPort)
}

//redis set 命令
func (rp *RedisPool) RedisSet(key, data string) error {
	if rp.pool[0] != nil {
		c := rp.pool[0].Get()
		defer c.Close()
		_, err := c.Do("SET", key, data)
		return err
	}
	return errors.New("not new redis pool!")
}

//redis get 命令
func (rp *RedisPool) RedisGet(key string) (string, error) {
	if rp.pool[0] != nil {
		c := rp.pool[0].Get()
		defer c.Close()
		res, err := redis.String(c.Do("GET", key))
		if err == nil {
			return res, err
		}
	}
	return "", errors.New("not new redis pool!")
}

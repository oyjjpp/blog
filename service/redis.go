package service

import (
	"errors"
	"strings"
	"time"

	"github.com/oyjjpp/blog/util/lib"

	"github.com/gomodule/redigo/redis"
)

//客户端实现redis连接池
type RedisPool struct {
	//连接池
	pool     []*redis.Pool
	pool_num int

	//一致性hash结构
	consistent_hash *lib.Consistent
}

//初始化redis连接池
func (rp *RedisPool) InitRedisPool(ipPort []string) {
	rp.pool = make([]*redis.Pool, 0, 30)

	rp.consistent_hash = lib.NewConsistent()

	for i := 0; i < len(ipPort); i++ {
		//拆分链接配置
		//127.0.0.1:6379@auth
		ipInfo := strings.Split(ipPort[i], "@")
		rPool := &redis.Pool{
			//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
			MaxIdle: 1024,
			//最大的激活连接数，表示同时最多有N个连接
			MaxActive: 2048,
			//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
			IdleTimeout: 180 * time.Second,
			//创建链接
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
		//分布redis节点
		rp.consistent_hash.Add(lib.NewNode(i, ipPort[i], 1))
	}
	rp.pool_num = len(ipPort)
}

//redis set 命令
func (rp *RedisPool) RedisSet(key, data string) error {
	keyHash := rp.consistent_hash.Get(key)
	if rp.pool[keyHash.Id] != nil {
		c := rp.pool[keyHash.Id].Get()
		defer c.Close()
		_, err := c.Do("SET", key, data)
		return err
	}
	return errors.New("not new redis pool!")
}

//redis get 命令
func (rp *RedisPool) RedisGet(key string) (string, error) {
	keyHash := rp.consistent_hash.Get(key)
	if rp.pool[keyHash.Id] != nil {
		c := rp.pool[keyHash.Id].Get()
		defer c.Close()
		res, err := redis.String(c.Do("GET", key))
		if err == nil {
			return res, err
		}
	}
	return "", errors.New("not new redis pool!")
}

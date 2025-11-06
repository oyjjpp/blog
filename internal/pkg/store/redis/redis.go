package redis

import (
	"context"

	redisdb "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	RedisClient *redisdb.Client // Redis 数据库
)

func Init() error {
	RedisClient = redisdb.NewClient(&redisdb.Options{
		Addr:     viper.GetString("redis.arrd"),
		Password: viper.GetString("redis.password"),
		PoolSize: viper.GetInt("redis.poolsize"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	// xlog.Info("init redis client", RedisClient)
	if err != nil {
		// xlog.Panic(err)
	}
	return err
}

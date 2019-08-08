package redisdb

import (
	"github.com/go-redis/redis"
	"github.com/phjt-go/logger"
	"github.com/spf13/viper"
)

// RedisDB Redis的DB对象
var RedisDB *redis.Client

func Init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis_host"),
		Password: viper.GetString("redis_pass"),
		DB:       viper.GetInt("redis_db"),
	})

	defer func() {
		if r := recover(); r != nil {
			logger.Error("Redis connection error,", r)
		}
	}()
	_, err := RedisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
	logger.Info("Redis connection ok")
}

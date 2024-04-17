package myRedis

import (
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Db       int
}

func GetRedisClient(config RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Db,
	})
	return rdb
}

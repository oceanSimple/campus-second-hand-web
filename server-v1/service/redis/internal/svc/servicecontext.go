package svc

import (
	"github.com/redis/go-redis/v9"
	"server-v1/internal/myRedis"
	"server-v1/service/redis/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: myRedis.GetRedisClient(myRedis.RedisConfig{
			Host:     c.RedisConfig.Host,
			Port:     c.RedisConfig.Port,
			Password: c.RedisConfig.Password,
			Db:       c.RedisConfig.Db,
		}),
	}
}

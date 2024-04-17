package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	RedisConfig struct {
		Host     string
		Port     int
		Password string
		Db       int
	}
}

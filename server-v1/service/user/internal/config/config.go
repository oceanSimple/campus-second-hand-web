package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		Username string
		Password string
		Host     string
		Port     int
		Dbname   string
	}
}

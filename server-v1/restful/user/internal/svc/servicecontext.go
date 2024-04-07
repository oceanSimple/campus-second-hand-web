package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"server-v1/restful/user/internal/config"
	"server-v1/service/user/accountservice"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc accountservice.AccountService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: accountservice.NewAccountService(zrpc.MustNewClient(c.UserRpc)),
	}
}

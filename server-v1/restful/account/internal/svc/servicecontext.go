package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"server-v1/restful/account/internal/config"
	"server-v1/service/account/accountservice"
)

type ServiceContext struct {
	Config     config.Config
	AccountRpc accountservice.AccountService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AccountRpc: accountservice.NewAccountService(zrpc.MustNewClient(c.AccountRpc)),
	}
}

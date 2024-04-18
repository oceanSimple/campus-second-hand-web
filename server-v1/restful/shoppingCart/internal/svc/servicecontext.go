package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"server-v1/restful/shoppingCart/internal/config"
	"server-v1/service/redis/client/shoppingcartservice"
)

type ServiceContext struct {
	Config          config.Config
	ShoppingCartRpc shoppingcartservice.ShoppingCartService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ShoppingCartRpc: shoppingcartservice.
			NewShoppingCartService(zrpc.MustNewClient(c.RedisRpc)),
	}
}

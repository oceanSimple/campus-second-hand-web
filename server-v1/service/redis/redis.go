package main

import (
	"flag"
	"fmt"

	"server-v1/service/redis/internal/config"
	shoppingcartserviceServer "server-v1/service/redis/internal/server/shoppingcartservice"
	"server-v1/service/redis/internal/svc"
	"server-v1/service/redis/pb/redisService"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/myRedis.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		redisService.RegisterShoppingCartServiceServer(grpcServer, shoppingcartserviceServer.NewShoppingCartServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

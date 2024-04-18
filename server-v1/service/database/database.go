package main

import (
	"flag"
	"fmt"
	"server-v1/service/database/internal/config"
	accountserviceServer "server-v1/service/database/internal/server/accountservice"
	orderserviceServer "server-v1/service/database/internal/server/orderservice"
	productserviceServer "server-v1/service/database/internal/server/productservice"
	"server-v1/service/database/internal/svc"
	"server-v1/service/database/pb/databaseService"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/database.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		databaseService.RegisterAccountServiceServer(grpcServer, accountserviceServer.NewAccountServiceServer(ctx))
		databaseService.RegisterProductServiceServer(grpcServer, productserviceServer.NewProductServiceServer(ctx))
		databaseService.RegisterOrderServiceServer(grpcServer, orderserviceServer.NewOrderServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

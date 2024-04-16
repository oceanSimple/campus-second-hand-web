// Code generated by goctl. DO NOT EDIT.
// Source: database.proto

package productservice

import (
	"context"

	"server-v1/service/database/pb/databaseService"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Account             = databaseService.Account
	AccountRequest      = databaseService.AccountRequest
	BoolResp            = databaseService.BoolResp
	IdRequest           = databaseService.IdRequest
	Order               = databaseService.Order
	Product             = databaseService.Product
	UpdateIntRequest    = databaseService.UpdateIntRequest
	UpdateStringRequest = databaseService.UpdateStringRequest

	ProductService interface {
		// 根据id获取商品信息
		GetProductById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Product, error)
	}

	defaultProductService struct {
		cli zrpc.Client
	}
)

func NewProductService(cli zrpc.Client) ProductService {
	return &defaultProductService{
		cli: cli,
	}
}

// 根据id获取商品信息
func (m *defaultProductService) GetProductById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Product, error) {
	client := databaseService.NewProductServiceClient(m.cli.Conn())
	return client.GetProductById(ctx, in, opts...)
}
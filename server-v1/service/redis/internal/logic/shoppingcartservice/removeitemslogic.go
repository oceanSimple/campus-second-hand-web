package shoppingcartservicelogic

import (
	"context"
	"server-v1/service/redis/internal/svc"
	"server-v1/service/redis/pb/redisService"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveItemsLogic {
	return &RemoveItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveItemsLogic) RemoveItems(in *redisService.RemoveItemsRequest) (*redisService.ResultResponse, error) {
	var rmLogic = RemoveItemLogic{}
	var result uint64
	var err error
	for i := 0; i < len(in.ProductIds); i++ {
		tempResult, err := rmLogic.RemoveItem(&redisService.RemoveItemRequest{
			UserId:    in.UserId,
			ProductId: in.ProductIds[i],
		})
		if err != nil {
			return nil, err
		}
		result += tempResult.Result
	}
	return &redisService.ResultResponse{
		Result: result,
	}, err
}

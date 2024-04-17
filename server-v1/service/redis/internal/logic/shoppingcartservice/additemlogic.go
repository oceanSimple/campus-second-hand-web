package shoppingcartservicelogic

import (
	"context"
	"strconv"

	"server-v1/service/redis/internal/svc"
	"server-v1/service/redis/pb/redisService"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddItemLogic) AddItem(in *redisService.AddItemRequest) (*redisService.ResultResponse, error) {
	var rdb = l.svcCtx.RedisClient
	result, err := rdb.HSet(l.ctx, strconv.FormatUint(in.UserId, 10), in.ProductId, in.Item).Result()
	return &redisService.ResultResponse{
		Result: uint64(result),
	}, err
}

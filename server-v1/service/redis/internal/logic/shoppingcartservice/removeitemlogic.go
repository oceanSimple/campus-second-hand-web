package shoppingcartservicelogic

import (
	"context"
	"strconv"

	"server-v1/service/redis/internal/svc"
	"server-v1/service/redis/pb/redisService"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveItemLogic {
	return &RemoveItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveItemLogic) RemoveItem(in *redisService.RemoveItemRequest) (*redisService.ResultResponse, error) {
	var rdb = l.svcCtx.RedisClient
	result, err := rdb.HDel(l.ctx, strconv.FormatUint(in.GetUserId(), 10), strconv.FormatUint(in.GetProductId(), 10)).Result()
	return &redisService.ResultResponse{
		Result: uint64(result),
	}, err
}

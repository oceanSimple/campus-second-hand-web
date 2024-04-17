package shoppingcartservicelogic

import (
	"context"
	"strconv"

	"server-v1/service/redis/internal/svc"
	"server-v1/service/redis/pb/redisService"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemsLogic {
	return &GetItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetItemsLogic) GetItems(in *redisService.IdRequest) (*redisService.ItemsResponse, error) {
	var rdb = l.svcCtx.RedisClient
	var m, err = rdb.HGetAll(l.ctx, strconv.FormatUint(in.Id, 10)).Result()
	if err != nil {
		return nil, err
	}
	var data = make([]string, 0, len(m))
	for _, v := range m {
		data = append(data, v)
	}
	return &redisService.ItemsResponse{
		Items: data,
	}, nil
}

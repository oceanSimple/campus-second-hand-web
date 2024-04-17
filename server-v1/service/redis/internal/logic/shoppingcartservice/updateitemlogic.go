package shoppingcartservicelogic

import (
	"context"

	"server-v1/service/redis/internal/svc"
	"server-v1/service/redis/pb/redisService"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemLogic {
	return &UpdateItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateItemLogic) UpdateItem(in *redisService.AddItemRequest) (*redisService.ResultResponse, error) {
	// 修改和添加逻辑一样
	logic := AddItemLogic{}
	return logic.AddItem(in)
}

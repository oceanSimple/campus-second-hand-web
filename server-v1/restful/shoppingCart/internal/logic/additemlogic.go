package logic

import (
	"context"
	"server-v1/service/redis/client/shoppingcartservice"

	"server-v1/restful/shoppingCart/internal/svc"
	"server-v1/restful/shoppingCart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddItemLogic) AddItem(req *types.AddReq) (resp *types.CountResp, err error) {
	result, err := l.svcCtx.ShoppingCartRpc.AddItem(l.ctx, &shoppingcartservice.AddItemRequest{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Item:      req.ShoppingCartItem,
	})
	if err != nil {
		return nil, err
	}
	return &types.CountResp{
		Count: result.Result,
	}, nil
}

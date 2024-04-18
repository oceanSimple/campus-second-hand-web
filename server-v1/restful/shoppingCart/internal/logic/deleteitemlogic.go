package logic

import (
	"context"
	"server-v1/service/redis/client/shoppingcartservice"

	"server-v1/restful/shoppingCart/internal/svc"
	"server-v1/restful/shoppingCart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteItemLogic {
	return &DeleteItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteItemLogic) DeleteItem(req *types.DeleteItemReq) (resp *types.CountResp, err error) {
	result, err := l.svcCtx.ShoppingCartRpc.RemoveItem(l.ctx, &shoppingcartservice.RemoveItemRequest{
		UserId:    req.UserId,
		ProductId: req.ProductId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CountResp{
		Count: result.Result,
	}, nil
}

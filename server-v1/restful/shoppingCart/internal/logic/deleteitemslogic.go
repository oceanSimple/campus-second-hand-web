package logic

import (
	"context"
	"server-v1/service/redis/client/shoppingcartservice"

	"server-v1/restful/shoppingCart/internal/svc"
	"server-v1/restful/shoppingCart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteItemsLogic {
	return &DeleteItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteItemsLogic) DeleteItems(req *types.DeleteItemsReq) (resp *types.CountResp, err error) {
	result, err := l.svcCtx.ShoppingCartRpc.RemoveItems(l.ctx, &shoppingcartservice.RemoveItemsRequest{
		UserId:     req.UserId,
		ProductIds: req.ProductId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CountResp{
		Count: result.Result,
	}, nil
}

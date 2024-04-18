package logic

import (
	"context"
	"server-v1/restful/shoppingCart/internal/svc"
	"server-v1/restful/shoppingCart/internal/types"
	"server-v1/service/redis/client/shoppingcartservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemsLogic {
	return &GetItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetItemsLogic) GetItems(req *types.IdReq) (resp *types.GetItemsResp, err error) {
	data, err := l.svcCtx.ShoppingCartRpc.GetItems(l.ctx, &shoppingcartservice.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	// 交给前端处理
	//var items = make([]model.ShoppingCartItem, 0, len(data.Items))
	//// 将json字符串转换为结构体
	//for _, itemJson := range data.Items {
	//	var item model.ShoppingCartItem
	//	_ = json.Unmarshal([]byte(itemJson), &item)
	//	items = append(items, item)
	//}
	//// 按照添加时间排序
	//sort.Slice(items, func(i, j int) bool {
	//	return items[i].Time.Sub(items[j].Time) > 0
	//})
	return &types.GetItemsResp{
		Items: data.Items,
	}, nil
}

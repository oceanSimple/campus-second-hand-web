package orderservicelogic

import (
	"context"
	"server-v1/internal/model"

	"server-v1/service/database/internal/svc"
	"server-v1/service/database/pb/databaseService"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据id获取订单信息
func (l *GetOrderByIdLogic) GetOrderById(in *databaseService.IdRequest) (*databaseService.Order, error) {
	var order model.HistoryOrder
	var result = l.svcCtx.Db.Where("is_deleted != ?", 1).
		Where("id = ?", in.Id).
		First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &databaseService.Order{
		Id:        order.Id,
		ProductId: order.ProductId,
		BuyerId:   order.BuyerId,
		SellerId:  order.SellerId,
		Price:     order.Price,
		Title:     order.Title,
	}, nil
}

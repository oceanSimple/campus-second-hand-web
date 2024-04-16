package productservicelogic

import (
	"context"
	"server-v1/internal/model"

	"server-v1/service/database/internal/svc"
	"server-v1/service/database/pb/databaseService"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByIdLogic {
	return &GetProductByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据id获取商品信息
func (l *GetProductByIdLogic) GetProductById(in *databaseService.IdRequest) (*databaseService.Product, error) {
	var product model.Product
	var result = l.svcCtx.Db.Where("is_deleted != ?", 1).
		Where("id = ?", in.Id).
		First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &databaseService.Product{
		Id:          product.Id,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Stock:       uint64(product.Stock),
		Campus:      product.Campus,
		Address:     product.Address,
		SellerId:    product.SellerID,
		Picture1:    product.Picture1,
		Picture2:    product.Picture2,
		Picture3:    product.Picture3,
		Type1:       product.Type1,
		Type2:       product.Type2,
		Others:      product.Others,
	}, nil
}

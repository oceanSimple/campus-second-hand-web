package logic

import (
	"context"

	"server-v1/restful/shoppingCart/internal/svc"
	"server-v1/restful/shoppingCart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemLogic {
	return &UpdateItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateItemLogic) UpdateItem(req *types.AddReq) (resp *types.CountResp, err error) {
	logic := AddItemLogic{
		Logger: l.Logger,
		ctx:    l.ctx,
		svcCtx: l.svcCtx,
	}
	return logic.AddItem(req)
}

package logic

import (
	"context"
	"server-v1/internal/model"

	"server-v1/service/account/internal/svc"
	"server-v1/service/account/pb/accountDbService"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountStringLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAccountStringLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountStringLogic {
	return &UpdateAccountStringLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新string类型字段
func (l *UpdateAccountStringLogic) UpdateAccountString(in *accountDbService.UpdateStringRequest) (*accountDbService.BoolResp, error) {
	var result = l.svcCtx.Db.Model(&model.Account{}).
		Where("id = ?", in.Id).
		Update(in.Field, in.Value)
	if result.RowsAffected == 0 {
		return &accountDbService.BoolResp{Result: false}, result.Error
	} else {
		return &accountDbService.BoolResp{Result: true}, nil
	}
}

package logic

import (
	"context"
	"server-v1/internal/model"

	"server-v1/service/account/internal/svc"
	"server-v1/service/account/pb/accountDbService"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountIntLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAccountIntLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountIntLogic {
	return &UpdateAccountIntLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新int类型字段
func (l *UpdateAccountIntLogic) UpdateAccountInt(in *accountDbService.UpdateIntRequest) (*accountDbService.BoolResp, error) {
	var result = l.svcCtx.Db.Model(&model.Account{}).
		Where("id = ?", in.Id).
		Update(in.Field, in.Value)
	if result.RowsAffected == 0 {
		return &accountDbService.BoolResp{Result: false}, result.Error
	} else {
		return &accountDbService.BoolResp{Result: true}, nil
	}
}

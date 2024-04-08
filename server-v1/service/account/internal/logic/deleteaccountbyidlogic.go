package logic

import (
	"context"
	"server-v1/internal/model"

	"server-v1/service/account/internal/svc"
	"server-v1/service/account/pb/accountDbService"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAccountByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAccountByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAccountByIdLogic {
	return &DeleteAccountByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除账户
func (l *DeleteAccountByIdLogic) DeleteAccountById(in *accountDbService.GetByIdRequest) (*accountDbService.BoolResp, error) {
	var result = l.svcCtx.Db.Model(&model.Account{}).
		Where("id = ?", in.Id).
		Update("is_deleted", 1)
	if result.RowsAffected == 0 {
		return &accountDbService.BoolResp{Result: false}, result.Error
	} else {
		return &accountDbService.BoolResp{Result: true}, nil
	}
}

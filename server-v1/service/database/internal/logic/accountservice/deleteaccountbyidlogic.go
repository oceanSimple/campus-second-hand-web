package accountservicelogic

import (
	"context"
	"server-v1/internal/model"
	"server-v1/service/database/internal/svc"
	"server-v1/service/database/pb/databaseService"

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
func (l *DeleteAccountByIdLogic) DeleteAccountById(in *databaseService.IdRequest) (*databaseService.BoolResp, error) {
	var result = l.svcCtx.Db.Model(&model.Account{}).
		Where("id = ?", in.Id).
		Update("is_deleted", 1)
	if result.RowsAffected == 0 {
		return &databaseService.BoolResp{Result: false}, result.Error
	} else {
		return &databaseService.BoolResp{Result: true}, nil
	}
}

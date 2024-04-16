package accountservicelogic

import (
	"context"
	"server-v1/internal/model"
	"server-v1/service/database/internal/svc"
	"server-v1/service/database/pb/databaseService"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountByAccountLogic {
	return &GetAccountByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据账号获取账户信息
func (l *GetAccountByAccountLogic) GetAccountByAccount(in *databaseService.AccountRequest) (*databaseService.Account, error) {
	var account model.Account
	var result = l.svcCtx.Db.Where("is_deleted != ?", 1).
		Where("account = ?", in.Account).
		First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &databaseService.Account{
		Id:        account.Id,
		Account:   account.Account,
		Password:  account.Password,
		Phone:     account.Phone,
		Email:     account.Email,
		Nickname:  account.Nickname,
		Age:       uint32(account.Age),
		Gender:    uint32(account.Gender),
		Campus:    account.Campus,
		Dormitory: account.Dormitory,
	}, nil
}

package logic

import (
	"context"
	"server-v1/internal/model"
	"time"

	"server-v1/service/account/internal/svc"
	"server-v1/service/account/pb/accountDbService"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAccountLogic {
	return &CreateAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建账户
func (l *CreateAccountLogic) CreateAccount(in *accountDbService.Account) (*accountDbService.BoolResp, error) {
	var account = model.Account{
		Id:        in.Id,
		Account:   in.Account,
		Password:  in.Password,
		Phone:     in.Phone,
		Email:     in.Email,
		Nickname:  in.Nickname,
		Age:       uint8(in.Age),
		Gender:    uint8(in.Gender),
		Campus:    in.Campus,
		Dormitory: in.Dormitory,
		Common: model.Common{
			IsDeleted:   0,
			GmtCreate:   time.Now(),
			GmtModified: time.Now(),
		},
	}
	var result = l.svcCtx.Db.Create(&account)
	if result.Error != nil {
		return &accountDbService.BoolResp{Result: false}, result.Error
	} else {
		return &accountDbService.BoolResp{Result: true}, nil
	}
}

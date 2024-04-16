package accountservicelogic

import (
	"context"
	"server-v1/internal/model"
	"time"

	"server-v1/service/database/internal/svc"
	"server-v1/service/database/pb/databaseService"

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
func (l *CreateAccountLogic) CreateAccount(in *databaseService.Account) (*databaseService.BoolResp, error) {
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
		return &databaseService.BoolResp{Result: false}, result.Error
	} else {
		return &databaseService.BoolResp{Result: true}, nil
	}
}

package logic

import (
	"context"
	"server-v1/internal/model"

	"server-v1/service/user/internal/svc"
	"server-v1/service/user/pb/user_db_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountByIdLogic {
	return &GetAccountByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountByIdLogic) GetAccountById(in *user_db_service.GetAccountByIdRequest) (account *user_db_service.Account, err error) {
	// todo: add your logic here and delete this line
	var user model.Account
	l.svcCtx.Db.First(&user, 1)
	account = &user_db_service.Account{
		Id:       user.Id,
		Account:  user.Account,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	}
	return
}

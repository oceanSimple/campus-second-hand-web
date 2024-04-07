package logic

import (
	"context"
	"server-v1/service/user/accountservice"

	"server-v1/restful/user/internal/svc"
	"server-v1/restful/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() (resp *types.Account, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.UserRpc.GetAccountById(l.ctx, &accountservice.GetAccountByIdRequest{
		Id: 1,
	})
	resp = &types.Account{
		Id:   response.Id,
		Code: response.Account,
	}
	return
}

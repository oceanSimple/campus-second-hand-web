package logic

import (
	"context"
	"fmt"
	"server-v1/service/user/accountservice"

	"server-v1/restful/user/internal/svc"
	"server-v1/restful/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountLogic {
	return &GetAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountLogic) GetAccount(req *types.GetRequest) (resp *types.Account, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("req id : ", req.Id)
	response, err := l.svcCtx.UserRpc.GetAccountById(l.ctx, &accountservice.GetAccountByIdRequest{
		Id: 1,
	})
	resp = &types.Account{
		Id:   response.Id,
		Code: response.Account,
	}
	return
}

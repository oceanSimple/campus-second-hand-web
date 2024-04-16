// Code generated by goctl. DO NOT EDIT.
// Source: database.proto

package server

import (
	"context"

	"server-v1/service/database/internal/logic/accountservice"
	"server-v1/service/database/internal/svc"
	"server-v1/service/database/pb/databaseService"
)

type AccountServiceServer struct {
	svcCtx *svc.ServiceContext
	databaseService.UnimplementedAccountServiceServer
}

func NewAccountServiceServer(svcCtx *svc.ServiceContext) *AccountServiceServer {
	return &AccountServiceServer{
		svcCtx: svcCtx,
	}
}

// 根据id获取账户信息
func (s *AccountServiceServer) GetAccountById(ctx context.Context, in *databaseService.IdRequest) (*databaseService.Account, error) {
	l := accountservicelogic.NewGetAccountByIdLogic(ctx, s.svcCtx)
	return l.GetAccountById(in)
}

// 根据账号获取账户信息
func (s *AccountServiceServer) GetAccountByAccount(ctx context.Context, in *databaseService.AccountRequest) (*databaseService.Account, error) {
	l := accountservicelogic.NewGetAccountByAccountLogic(ctx, s.svcCtx)
	return l.GetAccountByAccount(in)
}

// 创建账户
func (s *AccountServiceServer) CreateAccount(ctx context.Context, in *databaseService.Account) (*databaseService.BoolResp, error) {
	l := accountservicelogic.NewCreateAccountLogic(ctx, s.svcCtx)
	return l.CreateAccount(in)
}

// 更新string类型字段
func (s *AccountServiceServer) UpdateAccountString(ctx context.Context, in *databaseService.UpdateStringRequest) (*databaseService.BoolResp, error) {
	l := accountservicelogic.NewUpdateAccountStringLogic(ctx, s.svcCtx)
	return l.UpdateAccountString(in)
}

// 更新int类型字段
func (s *AccountServiceServer) UpdateAccountInt(ctx context.Context, in *databaseService.UpdateIntRequest) (*databaseService.BoolResp, error) {
	l := accountservicelogic.NewUpdateAccountIntLogic(ctx, s.svcCtx)
	return l.UpdateAccountInt(in)
}

// 删除账户
func (s *AccountServiceServer) DeleteAccountById(ctx context.Context, in *databaseService.IdRequest) (*databaseService.BoolResp, error) {
	l := accountservicelogic.NewDeleteAccountByIdLogic(ctx, s.svcCtx)
	return l.DeleteAccountById(in)
}
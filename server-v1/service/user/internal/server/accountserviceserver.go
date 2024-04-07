// Code generated by goctl. DO NOT EDIT.
// Source: user_rpc.proto

package server

import (
	"context"

	"server-v1/service/user/internal/logic"
	"server-v1/service/user/internal/svc"
	"server-v1/service/user/pb/user_db_service"
)

type AccountServiceServer struct {
	svcCtx *svc.ServiceContext
	user_db_service.UnimplementedAccountServiceServer
}

func NewAccountServiceServer(svcCtx *svc.ServiceContext) *AccountServiceServer {
	return &AccountServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *AccountServiceServer) GetAccountById(ctx context.Context, in *user_db_service.GetAccountByIdRequest) (*user_db_service.Account, error) {
	l := logic.NewGetAccountByIdLogic(ctx, s.svcCtx)
	return l.GetAccountById(in)
}
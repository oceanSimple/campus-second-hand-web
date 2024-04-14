package logic

import (
	"context"

	"server-v1/restful/account/internal/svc"
	"server-v1/restful/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenLoginLogic {
	return &TokenLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLoginLogic) TokenLogin() (resp *types.BoolResp, err error) {
	return &types.BoolResp{
		Result: true,
	}, nil
}

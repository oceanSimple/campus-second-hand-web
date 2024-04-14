package logic

import (
	"context"
	"server-v1/internal/jwts"
	"server-v1/internal/tool"
	"server-v1/service/account/pb/accountDbService"

	"server-v1/restful/account/internal/svc"
	"server-v1/restful/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PswLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPswLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PswLoginLogic {
	return &PswLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PswLoginLogic) PswLogin(req *types.PasswordLoginReq) (resp *types.TokenResp, err error) {
	// 根据账号查询用户信息
	var account, queryErr = l.svcCtx.AccountRpc.GetAccountByAccount(l.ctx,
		&accountDbService.GetByAccountRequest{Account: "admin"})
	// 判断账号是否存在,并检查密码是否正确
	if queryErr != nil || account.Password != tool.GetSha256String(req.Password) {
		return &types.TokenResp{
			Token:  "",
			Result: false,
		}, queryErr
	}

	// 登录成功,生成token，并返回
	auth := l.svcCtx.Config.Auth
	token, jwtErr := jwts.GetToken(jwts.JwtPayLoad{
		Id:      account.Id,
		Account: account.Account,
	}, auth.AccessSecret, auth.AccessExpire)
	if jwtErr != nil {
		return &types.TokenResp{
			Token:  "",
			Result: false,
		}, jwtErr
	}
	return &types.TokenResp{
		Token:  token,
		Result: true,
	}, nil
}

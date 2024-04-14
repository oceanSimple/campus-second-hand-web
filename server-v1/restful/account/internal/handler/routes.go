// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"server-v1/restful/account/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/pswLogin",
				Handler: PswLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/account/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/tokenLogin",
				Handler: TokenLoginHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/account/v1"),
	)
}

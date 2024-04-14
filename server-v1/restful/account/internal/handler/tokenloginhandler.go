package handler

import (
	"net/http"
	"server-v1/internal/response"
	"server-v1/restful/account/internal/logic"
	"server-v1/restful/account/internal/svc"
)

func TokenLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTokenLoginLogic(r.Context(), svcCtx)
		resp, err := l.TokenLogin()
		response.HttpResult(r, w, resp, err)
	}
}

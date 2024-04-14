package handler

import (
	"net/http"
	"server-v1/internal/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server-v1/restful/account/internal/logic"
	"server-v1/restful/account/internal/svc"
	"server-v1/restful/account/internal/types"
)

func PswLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PasswordLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPswLoginLogic(r.Context(), svcCtx)
		resp, err := l.PswLogin(&req)
		response.HttpResult(r, w, resp, err)
	}
}

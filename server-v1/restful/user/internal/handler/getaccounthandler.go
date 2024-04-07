package handler

import (
	"net/http"
	"server-v1/internal/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server-v1/restful/user/internal/logic"
	"server-v1/restful/user/internal/svc"
	"server-v1/restful/user/internal/types"
)

func getAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetAccountLogic(r.Context(), svcCtx)
		resp, err := l.GetAccount(&req)
		response.HttpResult(r, w, resp, err)
	}
}

package handler

import (
	"net/http"
	"server-v1/internal/response"

	"server-v1/restful/user/internal/logic"
	"server-v1/restful/user/internal/svc"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test()
		response.HttpResult(r, w, resp, err)
	}
}

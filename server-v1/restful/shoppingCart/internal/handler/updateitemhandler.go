package handler

import (
	"net/http"
	"server-v1/internal/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server-v1/restful/shoppingCart/internal/logic"
	"server-v1/restful/shoppingCart/internal/svc"
	"server-v1/restful/shoppingCart/internal/types"
)

func UpdateItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateItemLogic(r.Context(), svcCtx)
		resp, err := l.UpdateItem(&req)
		response.HttpResult(r, w, resp, err)
	}
}

package response

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"server-v1/internal/myError"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	// 成功返回
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
		return
	}

	// 错误返回
	errCode := uint32(500)
	errMsg := "服务器开小差啦，稍后再来试一试"

	var e *myError.CodeError
	if errors.As(err, &e) {
		errCode = e.GetErrCode()
		errMsg = e.GetErrMsg()
	} else {
		errMsg = err.Error()
	}
	logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
	httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
}

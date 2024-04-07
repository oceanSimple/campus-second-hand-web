package response

type SuccessBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *SuccessBean {
	return &SuccessBean{200, "OK", data}
}

type ErrorBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode uint32, errMsg string) *ErrorBean {
	return &ErrorBean{errCode, errMsg}
}

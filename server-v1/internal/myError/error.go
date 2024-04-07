package myError

import "fmt"

type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 获取错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 获取错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

// Error 返回错误信息string
func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrorCode:%d，ErrorMsg:%s", e.errCode, e.errMsg)
}

// NewErrCodeMsg 创建错误码和错误信息
func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

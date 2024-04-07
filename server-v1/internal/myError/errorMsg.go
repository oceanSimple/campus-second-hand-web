package myError

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[CommonError] = "服务器出现了一些小问题，请稍后再试"
}

// GetErrMsg 获取错误信息
func GetErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器出现了一些小问题，请稍后再试"
	}
}

// IsCodeErr 判断错误码是否存在
func IsCodeErr(errCode uint32) bool {
	_, ok := message[errCode]
	return ok
}

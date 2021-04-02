package resp

type IErrorCode int

const (
	RESP_CODE_SUCCESS   IErrorCode = 0
	RESP_CODE_PARAM_ERR IErrorCode = 10001
)

func (code IErrorCode) GetPrompts() string {
	switch code {
	case RESP_CODE_SUCCESS:
		return ""
	case RESP_CODE_PARAM_ERR:
		return "请求参数错误"
	}
	return "unknown err"
}

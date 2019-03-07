package utils

// CommonError 百度小程序返回的通用错误json
type CommonError struct {
	ErrCode int64  `json:"errno"`
	ErrMsg  string `json:"error"`
	ErrDesc string `json:"error_description"`
}
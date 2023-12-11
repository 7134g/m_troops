package result

import "net/http"

type Response struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *Response {
	return &Response{http.StatusOK, http.StatusText(http.StatusOK), data}
}

func Error(code uint32, errMsg string) *Response {
	return &Response{Code: code, Msg: errMsg, Data: NullJson{}}
}

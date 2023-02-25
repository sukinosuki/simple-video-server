package common

type AppResponse[T any] struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Data   T      `json:"data"`
	ErrMsg string `json:"err_msg"`
}

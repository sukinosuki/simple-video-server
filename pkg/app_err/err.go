package app_err

type BusinessErr struct {
	Handler string `json:"handler"`
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
	Err     error  `json:"err"`
	ErrMsg  string `json:"err_msg"`
}

var (
	ServerErr = &BusinessErr{
		Msg:  "服务器错误",
		Code: 500,
	}
)

func (e *BusinessErr) Error() string {

	return e.Msg
}

func New(err error, handlerName string, msg string) *BusinessErr {

	return &BusinessErr{
		Handler: handlerName,
		Msg:     msg,
		Err:     err,
		Code:    500, // TODO
	}
}

func (e *BusinessErr) NewErr(err error, msg string) *BusinessErr {

	e.Err = err
	e.Msg = msg

	return e
}

func NewServerErr(err error, errMsg string) *BusinessErr {

	return &BusinessErr{
		Err:    err,
		ErrMsg: errMsg,
		Code:   ServerErr.Code,
		Msg:    ServerErr.Msg,
	}
}

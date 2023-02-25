package business_code

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"sync/atomic"
)

var (
	_message atomic.Value // NOTE:stored map[string]map[int]string
	//_codes   = map[int]struct{}{} //register codes
	_codes = map[int]string{} //register codes
)

type BusinessCodes interface {
	Error() string
	Code() int
	Message() string
	Equal(error) bool
}

type BusinessCode int

func (e BusinessCode) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e BusinessCode) Code() int {
	return int(e)
}

func (e BusinessCode) Message() string {
	if msg, ok := ECodeMap[e.Code()]; ok {
		return msg
	}

	return e.Error()
}

func (e BusinessCode) Equal(err error) bool {
	return EqualError(e, err)
}

func New(e int) BusinessCode {
	if e < 0 {
		panic("business code must greater than zero")
	}

	return add(e)
}

func add(e int) BusinessCode {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("code: %d already exist", e))
	}

	_codes[e] = ""

	return Int(e)
}

func Int(i int) BusinessCode {
	return BusinessCode(i)
}

func String(e string) BusinessCode {
	if e == "" {
		return OK
	}

	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}

	return BusinessCode(i)
}

func Cause(e error) BusinessCodes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(BusinessCodes)
	if ok {
		return ec
	}

	err := e.Error()

	return String(err)
}

func EqualError(code BusinessCodes, err error) bool {
	return Cause(err).Code() == code.Code()
}

func GetCodesMap() map[int]string {
	return _codes
}

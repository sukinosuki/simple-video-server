package common

type CodeValue[K int | string, T any] struct {
	Code  K
	Value T
}

type IntString = CodeValue[int, string]

type StringString = CodeValue[string, string]

func (cv *CodeValue[K, T]) Is(code K) bool {

	return cv.Code == code
}

type CodeValues[K int | string, T any] struct {
	Maps map[int]CodeValue[K, T]
}

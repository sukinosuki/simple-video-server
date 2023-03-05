package common

type CodeValue[T any] struct {
	Code  int
	Value T
}

func (cv *CodeValue[T]) Is(code int) bool {

	return cv.Code == code
}

type CodeValues[T any] struct {
	Maps map[int]CodeValue[T]
}

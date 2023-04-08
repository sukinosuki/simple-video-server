package common

import "encoding/json"

type CodeValue[K int | string, T any] struct {
	Code        K
	ValueString T
}

type IntString = CodeValue[int, string]

type StringString = CodeValue[string, string]

func (cv *CodeValue[K, T]) Is(code K) bool {

	return cv.Code == code
}

type CodeValues[K int | string, T any] struct {
	Maps map[int]CodeValue[K, T]
}

//func (cv *CodeValue[K, T]) MarshalJSON() ([]byte, error) {
//
//	bytes, err := json.Marshal(cv)
//
//	return bytes, err
//}

func (cv *CodeValue[K, T]) UnmarshalJSON(data []byte) error {

	var k K
	err := json.Unmarshal(data, &k)

	return err
}

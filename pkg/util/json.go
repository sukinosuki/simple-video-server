package util

import "encoding/json"

func ParseJson[T any](data string) (*T, error) {
	var t T
	err := json.Unmarshal([]byte(data), &t)

	return &t, err
}

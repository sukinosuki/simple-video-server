package main

import (
	"bytes"
	"fmt"
	"net/url"
)

func main() {
	values := url.Values{}
	values.Add("name", "hanami")
	values.Add("age", "26")

	fmt.Println("v ", values) // age=26&name=hanami

	fmt.Println("encode ", values.Encode()) // age=26&name=hanami

	buffer := bytes.NewBufferString(values.Encode())
	fmt.Println("buffer ", buffer) // age=26&name=hanami
}

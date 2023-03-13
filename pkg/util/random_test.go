package util

import (
	"fmt"
	"testing"
)

func TestRandomNumber(t *testing.T) {

	str := RandomNumber(5)

	fmt.Println("str ", str)
}

func TestRandomNumberString(t *testing.T) {
	for i := 1; i < 10; i++ {
		str := RandomNumberString(5)

		fmt.Println("str ", str)
	}
}

package util

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber(n int) string {

	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	num := rand.Intn(100000)         // 生成一个0到99999之间的随机整数
	str := fmt.Sprintf("%05d", num)  // 转换为字符串，并补零
	return str                       // 输出结果
}

func RandomNumberString(n int) string {
	seed := time.Now().UnixNano()

	rand.Seed(seed)                    // 设置随机数种子
	num := rand.Intn(100000 ^ n)       // 生成一个0到10^n-1之间的随机整数
	str := fmt.Sprintf("%0*d", n, num) // 转换为字符串，并补零
	return str                         // 返回结果
}

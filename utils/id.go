/**
 * @Auth: Nuts
 * @Date: 2021/3/5 2:13 下午
 */
package utils

import (
	"time"
)

var num2char = "0123456789abcdefghijklmnopqrstuvwxyz"

// 10进制转16或36进制
func NumToBHex(num, n int) string {
	numStr := ""
	for num != 0 {
		yu := num % n
		numStr = string(num2char[yu]) + numStr
		num = num / n
	}
	return numStr
}

// 唯一键生成函数
func GetUniqId() (uniqId string) {
	uniqId = NumToBHex(int(time.Now().UnixNano()), 36)
	return uniqId
}

// 唯一键生成函数
func GetUniqLenId(len int) (uniqId string) {
	uniqId = NumToBHex(int(time.Now().UnixNano()), len)
	return uniqId
}

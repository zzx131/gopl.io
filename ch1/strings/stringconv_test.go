package strings_learn

import (
	"fmt"
	"strconv"
	"testing"
)

// 将数字类型转string
func TestStringsConv(t *testing.T) {
	// var nam = "zhangZeXing" + 2 // 错误
	var nam = "zhangZeXing" + strconv.Itoa(3)
	fmt.Println(nam)
}

// 将字符串类型转成int类型
func TestStringConvTwo(t *testing.T) {
	var one = "12"
	result, _ := strconv.Atoi(one)
	fmt.Printf("%d\n", result)
}

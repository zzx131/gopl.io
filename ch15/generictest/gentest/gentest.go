package main

import (
	"errors"
	"fmt"
)

func show[num int64 | float64](s num) {
	fmt.Println(s)
}

// splitAnySlice any 使用
// 将数组切片切成两个数组切片
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// indexOf comparable 使用
// 数组寻找元素
func indexOf[T comparable](s []T, x T) (int, error) {
	for i, v := range s {
		if v == x {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func main() {
	var sum1 int64 = 28
	var sum2 float64 = 29.5
	show(sum1)
	show(sum2)
	// any的使用
	fmt.Println("==========any使用============")
	oneSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oneS, twoS := splitAnySlice(oneSlice)
	fmt.Println(oneS)
	fmt.Println(twoS)
	fmt.Println("==========使用结束===========")

	// comparable 的使用
	fmt.Println("==========comparable 使用============")
	res, err := indexOf(oneSlice, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("结果是：%d\n", res)
	}
	fmt.Println("==========comparable 使用结束=========")
}

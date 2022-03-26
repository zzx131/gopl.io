package for_learn

import (
	"fmt"
	"testing"
)

// 基于计数器的迭代
func TestForOne(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("this is the %d iteration\n", i)
	}
}

// 根据判断语句进行迭代
func TestForTwo(t *testing.T) {
	var i = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("the variable i is now:%d\n", i)
	}
}

// 用迭代器进行循环输出
func TestForThree(t *testing.T) {
	//oneSlice := make([]int, 10)
	oneSlice := []int{1, 2, 3, 4, 5}
	for _, value := range oneSlice {
		fmt.Printf("slice is value %d\n", value)
	}
}

// 用于迭代器进行无限循环
func TestForFore(t *testing.T) {
	flag := 1
	for true {
		if flag < 10 {
			fmt.Printf("value is %d\n", flag)
			flag++
		} else {
			break
		}
	}
	fmt.Println("this is end")
}

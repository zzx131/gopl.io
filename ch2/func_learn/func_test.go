package funclearn

import (
	"fmt"
	"testing"
)

// 寻找最小值
func TestFunc(t *testing.T) {
	x := min(1, 3, 2, 0)
	fmt.Printf("the minimum is: %d\n", x)
	slice := []int{7, 9, 5, 1, 2}
	x = min(slice...)
	fmt.Printf("the minium in the slice is:%d\n", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// defer 练习用于最后的资源释放，关键字用于最后输出，定义在函数前面
func TestDefer(t *testing.T) {
	fmt.Println("in function at the top")
	defer function1()
	fmt.Println("in function at the bottom")
}

func function1() {
	fmt.Println("the defer func")
}

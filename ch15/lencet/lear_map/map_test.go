package lear_slice

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSimple(t *testing.T) {
	// 初始化map
	// oneMap := map[string]int{"one": 1, "two": 2, "three": 3}
	oneMap := make(map[string]int, 10)
	oneMap["one"] = 1
	oneMap["two"] = 2
	oneMap["three"] = 3
	// 可以看出map是没有顺序的
	for _, v := range oneMap {
		fmt.Println("value:" + strconv.Itoa(v))
	}
	// 删除一个值,如果没有值，不会报错
	delete(oneMap, "one")
	fmt.Println(len(oneMap)) // 2
}

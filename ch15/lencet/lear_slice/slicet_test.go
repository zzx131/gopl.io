package lear_slice

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/slice"
)

// 判断是否包含
func TestContain(t *testing.T) {

	fmt.Println(slice.Contain([]string{"a", "b", "c"}, "a"))
	fmt.Println(slice.Contain([]string{"a", "b", "c"}, "d"))
	fmt.Println(slice.Contain([]string{""}, ""))
	fmt.Println(slice.Contain([]string{}, ""))
	fmt.Println(slice.Contain([]int{1, 2, 3}, 1))
}

// 判断是否包含子集合
func TestContainSub(t *testing.T) {
	oneSlice := []string{"zhangshan", "lishi", "wangwu"}
	twoSlice := []string{"wangwu1"}
	fmt.Println("containsub", slice.ContainSubSlice(oneSlice, twoSlice))
}

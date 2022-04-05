package lear_slice

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"golang.org/x/exp/maps"
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

func TestMapSort(t *testing.T) {
	// 将map进行排序
	oneMap := map[string]int{"a": 1, "c": 2, "d": 3, "b": 4}
	// 获取所有key
	mapKeys := maps.Keys(oneMap)
	// 将key进行排序
	sort.Strings(mapKeys)
	for _, keyv := range mapKeys {
		fmt.Println(oneMap[keyv])
	}
}

func TestNewMap(t *testing.T) {
	// oneMap := new(map[string]int)
	// oneMap["one"] = 1
}

package main

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

// indexOfInteger constraints 组合使用
func indexOfInteger[T constraints.Integer](s []T, x T) (int, error) {
	for i, v := range s {
		if v == x {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func main() {
	oneSlice := []int{1, 2, 3, 4}
	index, err := indexOfInteger(oneSlice, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", index)
}

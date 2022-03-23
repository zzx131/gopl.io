package gentest_test

import (
	"errors"
	"fmt"
	"testing"
)

func indexOfFromSli[T comparable](s []T, x T) (int, error) {
	for i, v := range s {
		if v == x {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}
func TestGen(t *testing.T) {
	oneSlice := []string{"张三", "李四", "王五"}
	result, err := indexOfFromSli(oneSlice, "李四")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", result)
}

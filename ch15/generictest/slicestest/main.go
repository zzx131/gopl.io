package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println(slices.Index([]string{"apple", "pear", "banana"}, "banana"))
}

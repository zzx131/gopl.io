package main

import "fmt"

// 练习go的默认转化，属于强制转化，
func main() {
	a := 5.0
	b := int(a)
	fmt.Printf("%d\n", b)
	// 将string转成int
	/*st := "12"
	stint := int(st)
	fmt.Printf("%d\n", stint)*/
}

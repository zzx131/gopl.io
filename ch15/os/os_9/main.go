package main

import (
	"fmt"
	"os"
	"strings"
)

// 从命令行读取参数(os.Args)
func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

// 从缓存区里面读取
func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was:%s\n", input)
	}
}

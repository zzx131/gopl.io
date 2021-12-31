package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCEFG")

	str := strings.NewReader("123455")

	// 带有buf的 read
	br := bufio.NewReader(s)

	b, _ := br.ReadString('\n')

	fmt.Println(b)

	// Reset丢弃缓冲中的数据，清除任何错误，将b重设为其下层从r读取数据。
	br.Reset(str)

	b, _ = br.ReadString('\n')

	fmt.Println(b)
}

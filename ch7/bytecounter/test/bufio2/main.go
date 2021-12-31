package main

import (
	"bufio"
	"fmt"
	"strings"
)

// 读取一个字节
func main() {
	// func (b *Reader) ReadByte() (c byte, err error)
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)

	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	// UnreadByte吐出最近一次读取操作读取的最后一个字节。（只能吐出最后一个，多次调用会出问题）
	br.UnreadByte()
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
}

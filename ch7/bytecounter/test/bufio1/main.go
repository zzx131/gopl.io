package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678")
	br := bufio.NewReader(s)

	p := make([]byte, 10)

	for {
		n, err := br.Read(p)

		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p):%v\n", string(p[0:n]))
		}
	}
}

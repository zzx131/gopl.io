// 7.1：使用类似ByteCounter的想法，实现单词和行的计数器，实现时考虑使用bufio.ScanWords。

package main

import (
	"bufio"
	"fmt"
)

type WordsCounter int

// Write 统计单词个数
func (c *WordsCounter) Write(content []byte) (int, error) {
	for start := 0; start < len(content); {
		advance, _, err := bufio.ScanWords(content[start:], true)

		if err != nil {
			return 0, err
		}

		start += advance

		(*c)++
	}

	return int(*c), nil
}

type LineCounter int

func (c *LineCounter) Write(content []byte) (int, error) {
	for start := 0; start < len(content); {
		advance, _, err := bufio.ScanLines(content[start:], true)

		if err != nil {
			return 0, err
		}

		start += advance

		(*c)++
	}
	return int(*c), nil
}

func main() {
	var wc WordsCounter
	wc.Write([]byte("Hello Worlds Test Me"))
	fmt.Println(wc) // 4

	wc.Write([]byte("append something to the end"))
	fmt.Println(wc) // 9

	var lc LineCounter
	fmt.Fprintf(&lc, "%s\n%s\n%s\n", "Hello Word", "Second Line", "Third Line")
	fmt.Println(lc) // 3

	fmt.Fprintf(&lc, "%s\n%s\n%s", "第4行", "第5行", "")
	fmt.Println(lc) // 5

}

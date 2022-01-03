// 7.2：实现一个满足如下签名的CountingWriter函数，输入一个io.Writer，
// 输出一个封装了输入值的心Writer，以及一个置项int64的指针，改制真对应的值是新的Writer吸入的字节数。
package main

import (
	"fmt"
	"io"
	"os"
)

// 记录数量
type CountWriter struct {
	Writer io.Writer
	Count  int
}

func (cw *CountWriter) Write(content []byte) (int, error) {
	n, err := cw.Writer.Write(content)

	if err != nil {
		return n, err
	}

	cw.Count += n

	return n, nil
}

func CountingWriter(write io.Writer) (io.Writer, *int) {
	cw := CountWriter{Writer: write}

	return &cw, &(cw.Count)
}

func main() {
	cw, counter := CountingWriter(os.Stdout)

	fmt.Fprintf(cw, "%s", "Print something to the screen ...")

	fmt.Println(*counter)

	cw.Write([]byte("Append something..."))

	fmt.Println(*counter)
}

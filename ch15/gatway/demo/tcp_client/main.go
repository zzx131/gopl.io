package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1,建立连接
	conn, err := net.Dial("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("connect failed, err:%v\n", err)
		return
	}
	defer conn.Close()
	// 2, 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 3、一直读取直到读到\n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v", err)
			break
		}
		// 4、读取Q时停止
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		// 5、回复服务器
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("weite failed, err:%v\n", err)
			break
		}
	}
}

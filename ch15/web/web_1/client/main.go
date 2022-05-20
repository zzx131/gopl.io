package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Tcp 请求
func main() {
	//打开连接:
	conn, err := net.Dial("tcp", "0.0.0.0:8070")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}
	defer conn.Close()
	// 2，读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
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

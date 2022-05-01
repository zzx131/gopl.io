package main

import (
	"fmt"
	"gopl.io/ch15/gatway/demo/unpack"
	"net"
)

func main() {
	// 1. 监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail,err:%v\n", err)
		return
	}
	// 2.接受请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail,err:%v\n", err)
			continue
		}
		// 3.创建协程
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect faile,err:%v\n", err)
			break
		}
		str := string(bt)
		fmt.Printf("receive from client, data:%v\n", str)
	}
}

package main

import (
	"fmt"
	"net"
)

func main() {
	// 1、监听tcp端口
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err %v", err)
		return
	}
	// 2、建立套接字
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail,err: %v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect failed, err:%v\n")
			break
		}
		str := string(buf[:n])
		fmt.Printf("receive from client, data:%v", str)
	}
}

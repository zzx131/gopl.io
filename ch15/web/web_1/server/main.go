package main

import (
	"fmt"
	"net"
)

// tcp 服务端
func main() {
	fmt.Println("Starting the server ...")
	// 1,创建 listener
	listener, err := net.Listen("tcp", "0.0.0.0:8070")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			continue
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
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

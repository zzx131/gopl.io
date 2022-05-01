package main

import (
	"fmt"
	"net"
)

func main() {
	// 1. 连接服务器
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("connect failed, err:%v\n", err)
		return
	}

	for i := 0; i < 100; i++ {
		// 2, 发送数据
		_, err := conn.Write([]byte("hello server!"))
		if err != nil {
			fmt.Printf("send data failed, err:%v\n", err)
			return
		}
		// 3. 接受数据
		reault := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(reault)
		if err != nil {
			fmt.Printf("receive data failed, err：%v\n", err)
			return
		}
		fmt.Printf("receive from addr:%v, data:%v\n", remoteAddr, string(reault[:n]))
	}
}

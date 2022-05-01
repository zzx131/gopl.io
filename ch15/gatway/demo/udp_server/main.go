package main

import (
	"fmt"
	"net"
)

// main server 代码
func main() {
	// 1.监听udp端口
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("listen fail, err:%v", err)
	}
	// 2.接受请求
	for {
		var data [1024]byte
		n, addr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("accept fail, err:%v\n", err)
			continue
		}
		// 3.创建协程,回复数据
		go func() {
			fmt.Printf("addr:%v, data:%v, count:%v\n", addr, string(data[:n]), n)
			_, err = listener.WriteToUDP([]byte("received success!"), addr)
			if err != nil {
				fmt.Printf("write failed, err:%v\n", err)
			}
		}()
	}
}

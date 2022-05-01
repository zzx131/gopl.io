package main

import (
	"fmt"
	"gopl.io/ch15/gatway/demo/unpack"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed,err:%v\n", err.Error())
		return
	}
	unpack.Encode(conn, "hello word 0!!!")
}

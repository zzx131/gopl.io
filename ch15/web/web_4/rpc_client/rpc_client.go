package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 同步调用， 乘法
	args := &Args{6, 3}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// 异步调用， 除法
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall, OK := <-divCall.Done // divCall.Done是一个channel，返回值是一个Call指针，并与divCall相等，结束前一直阻塞
	if divCall == replyCall && OK {
		fmt.Printf("Divide: %d/%d=%d, %d%%%d=%d\n", args.A, args.B, quotient.Quo, args.A, args.B, quotient.Rem)
	} else {
		log.Fatal("Divide error")
	}
}

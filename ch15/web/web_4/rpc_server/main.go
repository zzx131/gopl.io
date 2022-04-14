package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	time.Sleep(3 * time.Second)
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	time.Sleep(3 * time.Second)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234") //监听tcp:1234端口, 如果有多个网卡，冒号前也可以写IP
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}

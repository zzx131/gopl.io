package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

// 当所有线程停止的时候，会发生死锁
func main() {
	out := make(chan int)
	out <- 2
	go f1(out)
}

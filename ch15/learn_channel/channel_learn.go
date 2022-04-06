package main

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 2, "how many goroutines")

func f(left, right chan int) {
	left <- 1 + <-right
}
func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left chan int = nil
	var right chan int = leftmost

	for i := 0; i < *ngoroutine; i++ {
		left = right
		right = make(chan int)
		go f(left, right)
	}
	right <- 1      // bang!
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 10, ongeveer 1,5 s
}

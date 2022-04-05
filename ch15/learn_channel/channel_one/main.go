package main

import "fmt"

func main() {
	left := make(chan int)
	right := make(chan int)

	left <- 1 + <-right

	right <- 0

	fmt.Println(<-left)
}

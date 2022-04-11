package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	defer def()

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\n", e)
		}
	}()
	panic("A sever erroe occurred: stopping the program!")
	fmt.Println("Ending the program")
}

func def() {
	fmt.Println("def func...")
}

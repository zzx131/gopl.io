package main

import "io/ioutil"

func main() {
	ioutil.WriteFile("doc.txt", []byte("123456"), 0644)
}

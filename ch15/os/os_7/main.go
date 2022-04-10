package main

import "os"

// 不带缓存区进行写入
func main() {
	os.Stdout.WriteString("hello word\n")
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	f.WriteString("hello word in a file \n")
}

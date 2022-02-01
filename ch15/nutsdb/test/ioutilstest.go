package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	byt, _ := ioutil.ReadFile("/home/zzx/go-proj/GOPATH/src/gopl.io/ch15/nutsdb/test/doc.txt")
	result := strings.Split(string(byt), "\n")
	fmt.Println(len(result))
}

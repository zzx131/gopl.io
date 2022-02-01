package main

import (
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"io/ioutil"
	"strings"
)

func main() {
	byt, _ := ioutil.ReadFile("./doc.txt")
	result := strings.Split(string(byt), "\n")
	fmt.Println(fuzzy.Find("小明", result))
}

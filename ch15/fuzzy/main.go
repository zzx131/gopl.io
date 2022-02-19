package main

import (
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"io/ioutil"
	"strings"
)

func main() {
	byt, _ := ioutil.ReadFile("./doc.txt")
	stringArray := make([]string, 0)
	stringArray = append(stringArray, string(byt))

	fmt.Println(strings.Contains(string(byt), "小明"))
	fmt.Println(fuzzy.Find("小明 ", stringArray))
}

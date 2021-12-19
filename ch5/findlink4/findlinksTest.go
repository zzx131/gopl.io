package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
)

/**
练习 5.1： 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。

练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。

练习 5.3： 编写函数输出所有text结点的内容。注意不要访问

练习 5.4： 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
*/
func main() {
	url := "http://goland.org"
	fmt.Println(url)
	//fetch返回字节类型，使用字节转流，被Parse读取
	doc, err := html.Parse(bytes.NewReader(fetch(url)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "err:%v\n", err)
		os.Exit(1)
	}
	// 打印所有链接
	for _, link := range visitTest(nil, doc) {
		fmt.Println(link)
	}
	// 出现的标签
	outline(nil, doc)
	fmt.Println("统计各个标签出现的数量")
	var m = make(map[string]int)
	m = countH(m, doc)
	for k, v := range m {
		fmt.Printf("%s ===> %d\n", k, v)
	}
	// 输出 text 节点内容
	// outText(doc)
}

// 访问url
func fetch(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s:%v \n", url, err)
		os.Exit(1)
	}
	return b
}

// 遍历树节点
func visitTest(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visitTest(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visitTest(links, n.NextSibling)
	}
	return links
}

// 出现的标签栈
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		// 入栈
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 继续递归
		outline(stack, c)
	}
}

// 标签出现的次数
func countH(m map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m = countH(m, c)
	}
	return m
}

// 输出 text节点
func outText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	// 子节点
	if c := n.FirstChild; c != nil && c.Data != "style" && c.Data != "script" {
		outText(c)
	}
	// 同级下的标签
	if c := n.NextSibling; c != nil && c.Data != "style" && c.Data != "script" {
		outText(c)
	}
}

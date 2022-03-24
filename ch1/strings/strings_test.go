package strings_learn

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

// 判断是否有前缀或者是后缀
func TestStringsOne(t *testing.T) {
	var str = "This is an example string 张泽新"
	// 判断是否有前缀
	fmt.Println(strings.HasPrefix(str, "Th"))
	// 判断是否有后缀
	fmt.Println(strings.HasSuffix(str, "ing"))
	// 字符串包含关系
	fmt.Println(strings.Contains(str, "isz"))
	fmt.Println(strings.ContainsAny(str, "isan"))
	fmt.Println(strings.ContainsRune(str, '张'))
}

// 注意golang中一个汉字占3个字节
func TestRune(t *testing.T) {
	var chinese = "我是中国人，i am chinese"
	fmt.Println("chinese length", len(chinese))
	fmt.Println("chines word length", len([]rune(chinese)))
	fmt.Println("chines word length", utf8.RuneCountInString(chinese))
}

// 判断子字符串或字符在父字符串中出现的位置（索引）
func TestStringIndex(t *testing.T) {
	var str = "Hi, i im Marc hi"
	fmt.Printf("%d\n", strings.Index(str, "Marc"))
}

// 将切片拼接成一个字符串
func TestStringsJoin(t *testing.T) {
	// oneslicep := make([]int, 0)
	oneslicep := []string{"1", "2", "3", "4", "5"}
	fmt.Println(strings.Join(oneslicep, ","))
}

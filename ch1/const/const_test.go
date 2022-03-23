package lear_const

import (
	"fmt"
	"testing"
)

// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
const (
	a = iota
	b
	c
	d = 5
	e
	f
)

// 2*2 的常量矩阵，iota 只会增长一次，而不会因为使用了两次就增长两次
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, fig
)

// 使用 iota 结合 位运算 表示资源状态的使用案例
const (
	Open    = 1 << iota // 0001
	Close               // 0010
	Pending             // 0100
)

// 可以使用下划线进行
const (
	_  = iota             // 使用 _ 忽略不需要的 iota
	KB = 1 << (10 * iota) // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
	TB                    // 1 << (10*4)
	PB                    // 1 << (10*5)
	EB                    // 1 << (10*6)
	ZB                    // 1 << (10*7)
	YB                    // 1 << (10*8)
)

const (
	one = iota
	_
	three
)

func TestConstOne(t *testing.T) {
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", d)
	fmt.Printf("%d\n", e)
	fmt.Printf("%d\n", f)
}

func TestConstTwo(t *testing.T) {
	fmt.Printf("%d\n", Apple)      // 1
	fmt.Printf("%d\n", Banana)     // 2
	fmt.Printf("%d\n", Cherimoya)  // 2
	fmt.Printf("%d\n", Durian)     // 3
	fmt.Printf("%d\n", Elderberry) // 3
	fmt.Printf("%d\n", fig)        // 4
}

func TestConstThree(t *testing.T) {
	fmt.Println(Open)
	fmt.Println(Close)
	fmt.Println(Pending)
}

func TestConstFor(t *testing.T) {
	fmt.Printf("one=%d , three=%d", one, three)
}

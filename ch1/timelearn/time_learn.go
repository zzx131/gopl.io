package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	// 将字符串转成时间
	time, _ := datetime.FormatStrToTime("2006-01-02 15:04:05", "yyyy-mm-dd hh:mm:ss")
	fmt.Println("将字符转成时间", time)
	// 将时间转成字符串
	result := datetime.FormatTimeToStr(t, "yyyy-mm-dd")
	fmt.Println(result)
}

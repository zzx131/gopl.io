package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// 解析传入的参数
	flag.Parse()

	defer glog.Flush()

	glog.Info("This is info log")

	glog.Warning("This is a warning log")

	glog.Error("This is a weeor log")

	glog.Fatal("This is a fatal log")

}

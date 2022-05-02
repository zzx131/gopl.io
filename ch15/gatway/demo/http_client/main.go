package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type UserData struct {
	Id   int32
	Name string
}

func main() {
	// 创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // 试探时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  // 100-continue状态码超时时间
	}
	// 创建客户端连接
	client := &http.Client{
		Timeout:   time.Second * 30,
		Transport: transport,
	}
	// 发送Post请求
	var reqData = &UserData{
		Id:   1,
		Name: "zhangsan",
	}
	byteReq, _ := json.Marshal(reqData)
	resp, err := client.Post("http://127.0.0.1:1210/bye", "application/json", bytes.NewReader(byteReq))

	// 发送get请求
	// resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 读取内容
	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bds))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/bye", sayBye)
	// 创建服务器
	server := &http.Server{
		Addr:         ":1210",
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	// 监听端口并提供服务
	log.Fatal(server.ListenAndServe())
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(1 * time.Second)
	var requestMap = make(map[string]any)
	err := json.NewDecoder(r.Body).Decode(&requestMap)
	if err != nil {
		fmt.Printf("json decode err %v", err)
	}
	w.Write([]byte("bye bye, this is httpServer"))
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	// 解析url传参
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form{
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	// 写到客户端
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./login.html")
		log.Println(t.Execute(w, nil))
	} else {
		// 请求的是登录数据，那么执行登录逻辑判断
		r.ParseForm()
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func main()  {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil)

	if err != nil{
		log.Fatal("ListenAndServe", err)
	}
}

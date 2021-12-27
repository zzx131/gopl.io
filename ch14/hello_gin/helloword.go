package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// hello word
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello word")
	})

	// restful 路由，: 表达式
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello %s", name)
	})

	// gin 还提供了 * 号处理参数。
	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})

	// query string
	router.GET("/welocome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname")

		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8000")
}

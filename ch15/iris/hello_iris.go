package main

import (
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func main() {
	app := iris.New()
	accesslog(app)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})

	app.Handle("GET", "/contact", func(context iris.Context) {
		context.HTML("<h1>hello form /contact</h1>")
	})
	app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))
	// 通过YAML配置文件
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
}

func accesslog(app *iris.Application) {
	path := "iris"
	writer, _ := rotatelogs.New(
		path+"-%Y-%m-%d.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	// 日志写入文件
	app.Logger().SetOutput(writer)
	// 日志同时写入控制台，如果不想显示控制台可注释此语句
	app.Logger().AddOutput(os.Stdout)
	// 记录路由日志
	app.Use(logger.New(logger.Config{
		Status:     true,
		IP:         true,
		Method:     true,
		Path:       true,
		Query:      true,
		LogFunc:    nil,
		LogFuncCtx: nil,
		Skippers:   nil,
	}))
}

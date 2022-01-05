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

	path := "iris"

	writer, _ := rotatelogs.New(
		path+"-%Y-%m-%d.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),

		//这里设置1分钟产生一个日志文件
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	app.Logger().SetOutput(writer)    //日志写入文件
	app.Logger().AddOutput(os.Stdout) //日志同时写入控制台，如果不想显示控制台可注释此语句

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

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})

	app.Handle("GET", "/contact", func(context iris.Context) {
		context.HTML("<h1>hello form /contact</h1>")
	})
	// [...]
	//我们可以用这种方法单独定义我们的配置项
	app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))
	//也可以使用app.run的第二个参数
	/*app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))*/
	// 通过YAML配置文件
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
	//通过多参数配置 但是上面两种方式是我们最推荐的
	// 我们使用With+配置项名称 如WithCharset("UTF-8") 其中就是With+ Charset的组合
	//app.Run(iris.Addr(":8080"), iris.WithoutStartupLog, iris.WithCharset("UTF-8"))
	//当使用app.Configure(iris.WithoutStartupLog, iris.WithCharset("UTF-8"))设置配置项时
	//需要app.run()面前使用
}

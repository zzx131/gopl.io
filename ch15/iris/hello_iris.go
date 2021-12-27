package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]
	//我们可以用这种方法单独定义我们的配置项
	app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))
	//也可以使用app.run的第二个参数
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))
	//通过多参数配置 但是上面两种方式是我们最推荐的
	// 我们使用With+配置项名称 如WithCharset("UTF-8") 其中就是With+ Charset的组合
	//app.Run(iris.Addr(":8080"), iris.WithoutStartupLog, iris.WithCharset("UTF-8"))
	//当使用app.Configure(iris.WithoutStartupLog, iris.WithCharset("UTF-8"))设置配置项时
	//需要app.run()面前使用
}

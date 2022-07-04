package main

import (
	"log"

	"github.com/kataras/iris/v12"
)

// package main

// main
func main() {
	app := iris.Default()

	// 使用中间件
	app.Use(myMiddleware)

	// 定义接口
	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
		log.Println("aaa")
	})

	// 启动服务 默认0.0.0.0
	app.Run(iris.Addr(":8080"))
}

// myMiddleware
func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
}

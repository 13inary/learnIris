package main

import "github.com/kataras/iris/v12"

// package main

// main
func main() {
	app := iris.New()

	// 注册 静态文件
	app.RegisterView(iris.HTML("./views", ".html"))

	// 接口
	// 返回html 方法1
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "hello world!") // html里面的变量 --- 变量的值
		ctx.View("hello.html")                  // 返回那个 html
	})
	// 返回html 方法2
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>hello world!</h1>")
	})
	// 获取url路径作为参数
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	// 运行服务
	app.Run(iris.Addr(":8080"))
}

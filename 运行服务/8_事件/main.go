package main

import "github.com/kataras/iris/v12"

// package main

// main
func main() {
	app := iris.Default()

	// 关闭服务时候触发(对所有主机)
	//app.ConfigureHost(func(h *iris.Supervisor) {
	//h.RegisterOnShutdown(func() {
	//println("server terminated")
	//})
	//})
	//app.Run(iris.Addr(":8080"))

	// 关闭服务时候触发(特定主机)
	app.Run(iris.Addr(":8080", func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			println("server terminated")
		})
	}))
}

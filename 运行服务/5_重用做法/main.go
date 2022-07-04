package main

import (
	"github.com/kataras/iris/v12"
	"github.com/valyala/tcplisten"
)

// package main

// main unix 和 bsd主机 可以使用重用端口的功能
func main() {
	app := iris.Default()

	listenerCfg := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
	}

	l, err := listenerCfg.NewListener("tcp", ":8080")
	if err != nil {
		app.Logger().Fatal(err)
	}

	app.Run(iris.Listener(l))
}

package main

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// package main

// main
func main() {
	app := iris.Default()

	h := app.NewHost(&http.Server{Addr: ":8080"})
	h.RegisterOnShutdown(func() {
		println("server terminated")
	})
	app.Run(iris.Raw(h.ListenAndServe))

}

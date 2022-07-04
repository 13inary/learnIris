package main

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// package main

// main
func main() {
	app := iris.Default()
	app.Run(iris.Server(&http.Server{Addr: ":8080"}))
}

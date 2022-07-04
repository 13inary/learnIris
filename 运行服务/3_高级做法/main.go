package main

import (
	"net"

	"github.com/kataras/iris/v12"
)

// package main

// main
func main() {
	app := iris.Default()

	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}
	app.Run(iris.Listener(l))
}

package main

import "github.com/kataras/iris/v12"

// package main

// main
func main() {
	app := iris.Default()
	app.Run(iris.Addr(":8080"))
}

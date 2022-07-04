package main

import "github.com/kataras/iris/v12"

// package main

// main
func main() {
	app := iris.Default()
	app.Run(iris.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))
	//通过https://letsencrypt.org免费提供的证书开启服务
	// app.Run(iris.AutoTLS(":443", "example.com", "admin@example.com"))
}

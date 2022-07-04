package main

import "github.com/kataras/iris/v12"

// package main

// main
func main() {
	app := iris.Default()

	app.Run(iris.Addr(":8080"), iris.WithTunneling)

	// 完整配置
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(
		iris.Configuration{
			Tunneling: iris.TunnelingConfiguration{
				AuthToken:    "my-ngrok-auth-client-token",
				Bin:          "/bin/path/for/ngrok",
				Region:       "eu",
				WebInterface: "127.0.0.1:4040",
				Tunnels: []iris.Tunnel{
					{
						Name: "myApp",
						Addr: ":8080",
					},
				},
			},
		}))

	// 返回公共域名
	//ctx.Application().ConfigurationReadOnly().GetVHost()
}

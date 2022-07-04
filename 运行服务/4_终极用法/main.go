package main

import (
	"net"
	"os"

	"github.com/kataras/iris/v12"
)

// package main

// main
func main() {
	app := iris.Default()

	// unix socket
	if errOs := os.Remove(socketFile); errOs != nil && !os.IsNotExist(errOs) {
		app.Logger().Fatal(errOs)
	}

	l, err := net.Listen("unix", socketFile)
	if err != nil {
		app.Logger().Fatal(err)
	}
	if err = os.Chmod(socketFile, mode); err != nil {
		app.Logger().Fatal(err)
	}

	app.Run(iris.Listener(l))
}

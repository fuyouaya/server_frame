package main

import (
	"server_frame/pkg/invoker"
	"server_frame/pkg/router"
)

func main() {
	invoker.Init()
	r := router.HttpServer()

	invoker.Log.Infof("运行端口号为：%d", 8080)
	r.Run()

}

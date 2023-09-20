package main

import (
	"server_frame/pkg/invoker"
	"server_frame/pkg/model/mysql"
	"server_frame/pkg/router"
	"server_frame/pkg/service"
)

func main() {
	invoker.Init()
	service.Init()

	mysql.Init(invoker.MainDB)

	r := router.HttpServer()
	invoker.Log.Infof("运行端口号为：%d", 8080)
	r.Run()
}

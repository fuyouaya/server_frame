package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	_ "server_frame/docs"
	"server_frame/pkg/core"
	"server_frame/pkg/invoker"
	v1 "server_frame/pkg/router/api/v1"
)

func HttpServer() *gin.Engine {
	r := invoker.Gin

	// 项目接口文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouter(r, "")

	return r
}

func apiRouter(c *gin.Engine, prefix string) {
	r := c.Group(prefix)
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello",
		})
	})

	{
		user := r.Group("/user")
		user.POST("/login", core.Handle(v1.UserLogin))
	}
}

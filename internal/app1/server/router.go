package server

import (
	"gin-layout/internal/app1/configs"
	"gin-layout/internal/app1/service"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(_ *configs.Bootstrap,
	demo *service.DemoService) *gin.Engine {
	srv := gin.Default()
	//健康检查
	srv.GET("/health_check", func(ctx *gin.Context) {
		//ctx.JSON(chttp.Success(ctx.Request.Context(), nil))
	})
	//v1版本
	version := srv.Group("/v1")

	interfaceRouter := version.Group("/demo")
	{
		interfaceRouter.GET("/get", demo.Get)
		interfaceRouter.GET("/list", demo.List)
	}

	return srv
}

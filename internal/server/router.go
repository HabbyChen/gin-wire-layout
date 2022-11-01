package server

import (
	"gin-layout/internal/service"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(greeterService *service.GreeterService) *gin.Engine {
	srv := gin.Default()
	srv.GET("/hello", greeterService.SayHello)
	return srv
}

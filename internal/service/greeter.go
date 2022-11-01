package service

import (
	"github.com/gin-gonic/gin"
)

// GreeterService is a greeter service.
type GreeterService struct {
	//Router *gin.Engine
}

func NewGreeterService() *GreeterService {
	s := &GreeterService{}
	//s.Router.GET("/hello", s.SayHello)
	return s
}

func (s *GreeterService) SayHello(ctx *gin.Context) {

	ctx.String(200, "Hello World")
}

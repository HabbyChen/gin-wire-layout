//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"gin-layout/internal/server"
	"gin-layout/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp() (*gin.Engine, func(), error) {
	panic(wire.Build(server.NewHTTPServer, service.NewGreeterService))
}

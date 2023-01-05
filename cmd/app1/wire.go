//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"gin-layout/internal/app1/biz"
	"gin-layout/internal/app1/configs"
	"gin-layout/internal/app1/repo"
	"gin-layout/internal/app1/server"
	"gin-layout/internal/app1/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(conf *configs.Bootstrap) (*gin.Engine, func(), error) {
	panic(wire.Build(server.ProviderServer, service.ProviderService, biz.ProviderBiz, repo.ProviderRepo))
}

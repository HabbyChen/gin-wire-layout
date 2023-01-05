// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gin-layout/internal/app1/biz"
	"gin-layout/internal/app1/configs"
	"gin-layout/internal/app1/repo"
	"gin-layout/internal/app1/server"
	"gin-layout/internal/app1/service"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(conf *configs.Bootstrap) (*gin.Engine, func(), error) {
	demoRepo := repo.NewRepoDemo()
	demoBiz := biz.NewDemoBiz(conf, demoRepo)
	demoService := service.NewDemoService(demoBiz)
	engine := server.NewHTTPServer(conf, demoService)
	return engine, func() {
	}, nil
}

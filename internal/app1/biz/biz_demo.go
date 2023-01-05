package biz

import (
	"gin-layout/internal/app1/configs"
	"github.com/gin-gonic/gin"
)

type DemoBiz struct {
	bootstrap *configs.Bootstrap
	demoRepo  DemoRepo
}
type DemoRepo interface {
	Add(ctx *gin.Context, info *DemoDo) (*DemoDo, error)
	Edit(ctx *gin.Context, info *DemoDo) (*DemoDo, error)
}
type DemoDo struct {
}

func NewDemoBiz(bootstrap *configs.Bootstrap, repo DemoRepo) *DemoBiz {
	return &DemoBiz{bootstrap: bootstrap, demoRepo: repo}
}

func (b DemoBiz) Get(_ *gin.Context, _ *DemoDo) (*DemoDo, error) {
	return nil, nil
}

func (b DemoBiz) List(_ *gin.Context, _ *DemoDo) (*DemoDo, error) {
	return nil, nil
}

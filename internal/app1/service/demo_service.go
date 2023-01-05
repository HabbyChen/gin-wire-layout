package service

import (
	"gin-layout/internal/app1/biz"
	"github.com/gin-gonic/gin"
)

type DemoService struct {
	DemoBiz *biz.DemoBiz
}

func NewDemoService(biz *biz.DemoBiz) *DemoService {
	s := &DemoService{}
	return s
}

func (s *DemoService) Get(ctx *gin.Context) {
	_, err := s.DemoBiz.Get(ctx, nil)
	if err != nil {
		ctx.JSON(200, err)
		return
	}
	ctx.String(200, "Hello World")
}

func (s *DemoService) List(ctx *gin.Context) {
	_, err := s.DemoBiz.List(ctx, nil)
	if err != nil {
		ctx.JSON(200, err)
		return
	}
	ctx.String(200, "Hello World")
}

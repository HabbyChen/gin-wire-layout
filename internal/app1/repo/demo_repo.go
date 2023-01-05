package repo

import (
	"gin-layout/internal/app1/biz"
	"github.com/gin-gonic/gin"
)

type RepoDemo struct {
}

func (receiver RepoDemo) Add(ctx *gin.Context, info *biz.DemoDo) (*biz.DemoDo, error) {
	//TODO implement me
	panic("implement me")
}

func (receiver RepoDemo) Edit(ctx *gin.Context, info *biz.DemoDo) (*biz.DemoDo, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepoDemo() biz.DemoRepo {
	return &RepoDemo{}
}

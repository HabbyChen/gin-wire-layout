package biz

import (
	"github.com/google/wire"
)

var ProviderBiz = wire.NewSet(NewDemoBiz)

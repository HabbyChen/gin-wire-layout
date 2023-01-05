package repo

import (
	"github.com/google/wire"
)

var ProviderRepo = wire.NewSet(NewRepoDemo)

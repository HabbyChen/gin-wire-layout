package service

import "github.com/google/wire"

// ProviderService is service providers.
var ProviderService = wire.NewSet(NewDemoService)

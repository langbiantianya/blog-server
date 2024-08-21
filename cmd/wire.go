package cmd

import "github.com/google/wire"

var Set = wire.NewSet(NewApiRoutes)

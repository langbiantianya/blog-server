package routes

import "github.com/google/wire"

var Set = wire.NewSet(NewEssayRouter, NewTagRouter)

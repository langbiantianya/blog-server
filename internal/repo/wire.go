package repo

import "github.com/google/wire"

var Set = wire.NewSet(NewEssayRepo, NewTagRepo, NewFileRepo)

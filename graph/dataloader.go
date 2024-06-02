package graph

import (
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/immrshc/go-gqlgent/graph/loader"
)

type Loaders struct {
	UserLoader dataloader.Interface
}

func NewLoaders() *Loaders {
	userLoader := &loader.UserLoader{}
	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader(
			userLoader.BatchGetUsers,
			dataloader.WithWait(time.Millisecond),
		),
	}
}

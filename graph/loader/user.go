package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/immrshc/go-gqlgent/graph/model"
)

type UserLoader struct{}

func (u *UserLoader) BatchGetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, len(keys))
	for idx, key := range keys {
		results[idx] = &dataloader.Result{
			Data: &model.User{
				ID:   key.String(),
				Name: "User: " + key.String(),
			},
			Error: nil,
		}
	}
	return results
}

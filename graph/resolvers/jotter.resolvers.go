package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/usecase-server/graph/generated"
	"github.com/vickywane/usecase-server/graph/model"
)

func (r *jotterResolver) Usecase(ctx context.Context, obj *model.Jotter) ([]*model.Usecase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *jotterResolver) Contributors(ctx context.Context, obj *model.Jotter) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *jotterResolver) CreatedBy(ctx context.Context, obj *model.Jotter) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Jotter returns generated.JotterResolver implementation.
func (r *Resolver) Jotter() generated.JotterResolver { return &jotterResolver{r} }

type jotterResolver struct{ *Resolver }

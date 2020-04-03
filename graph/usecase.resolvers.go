package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/vickywane/usecase-server/graph/generated"
	"github.com/vickywane/usecase-server/graph/model"
)

func (r *usecaseResolver) CreatedAt(ctx context.Context, obj *model.Usecase) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *usecaseResolver) UpdateAt(ctx context.Context, obj *model.Usecase) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Usecase returns generated.UsecaseResolver implementation.
func (r *Resolver) Usecase() generated.UsecaseResolver { return &usecaseResolver{r} }

type usecaseResolver struct{ *Resolver }

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

func (r *caseResolver) CreatedAt(ctx context.Context, obj *model.Case) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *caseResolver) UpdateAt(ctx context.Context, obj *model.Case) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Case returns generated.CaseResolver implementation.
func (r *Resolver) Case() generated.CaseResolver { return &caseResolver{r} }

type caseResolver struct{ *Resolver }

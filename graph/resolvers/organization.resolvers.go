package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/usecase-server/graph/generated"
	"github.com/vickywane/usecase-server/graph/model"
)

func (r *organizationResolver) CreatedBy(ctx context.Context, obj *model.Organization) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *organizationResolver) Cases(ctx context.Context, obj *model.Organization) ([]*model.Case, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *organizationResolver) Usecases(ctx context.Context, obj *model.Organization) ([]*model.Usecase, error) {
	panic(fmt.Errorf("not implemented"))
}

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

type organizationResolver struct{ *Resolver }

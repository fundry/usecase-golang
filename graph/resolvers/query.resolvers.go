package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/usecase-server/graph/generated"
	"github.com/vickywane/usecase-server/graph/model"
)

func (r *queryResolver) User(ctx context.Context, id int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Usecase(ctx context.Context, id int) (*model.Usecase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Usecases(ctx context.Context) ([]*model.Usecase, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Case(ctx context.Context, id int) (*model.Case, error) {
	// case is a reserved word in GoLang
	// causes a name conflict
	casea := model.Case{ID: id}

	if err := r.DB.Select(&casea); err != nil {
		return nil, err
	}

	return &casea, nil
}

func (r *queryResolver) Cases(ctx context.Context) ([]*model.Case, error) {
	var cases []*model.Case

	if err := r.DB.Model(&cases).Select(); err != nil {
		return nil, err
	}
	return cases, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

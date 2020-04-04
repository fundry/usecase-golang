package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/vickywane/usecase-server/graph/generated"
	"github.com/vickywane/usecase-server/graph/model"
)

func (r *mutationResolver) CreateCase(ctx context.Context, input model.NewCase) (*model.Case, error) {
	cases := model.Case{
		ID:        rand.Int(),
		Name:      input.Name,
		Author:    input.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &cases, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := r.DB.Insert(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) CreateUsecase(ctx context.Context, input model.NewUsecase) (*model.Usecase, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

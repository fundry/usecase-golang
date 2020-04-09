package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/vickywane/usecase-server/graph/generated"
	"github.com/vickywane/usecase-server/graph/model"
)

func (r *mutationResolver) CreateCase(ctx context.Context, input model.NewCase) (*model.Case, error) {
	cases := model.Case{
		ID:         rand.Int(),
		Name:       input.Name,
		Author:     input.Author,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Bookmarked: false,
	}

	if err := r.DB.Insert(&cases); err != nil {
		return nil, err
	}

	return &cases, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		Name:           input.Name,
		Email:          input.Email,
		CreatedAt:      time.Now(),
		IsOrganization: false,
	}

	if err := r.DB.Insert(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) CreateUsecase(ctx context.Context, input model.NewUsecase) (*model.Usecase, error) {
	usecase := model.Usecase{
		ID:         rand.Int(),
		Title:      input.Title,
		Author:     input.Author,
		Content:    input.Content,
		Completed:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Bookmarked: false,
	}

	if err := r.DB.Insert(&usecase); err != nil {
		return nil, err
	}
	return &usecase, nil
}

func (r *mutationResolver) CreateJotter(ctx context.Context, input model.NewJotter) (*model.Jotter, error) {
	jotter := model.Jotter{
		ID:           rand.Int(),
		Index:        input.Index,
		Name:         input.Name,
		Content:      input.Content,
		Contributors: input.Contributors,
		Completed:    false,
		Usecase:      input.Usecase,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := r.DB.Insert(&jotter); err != nil {
		return nil, err
	}

	return &jotter, nil
}

func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.NewOrganization) (*model.Organization, error) {
	organization := model.Organization{
		ID:        rand.Int(),
		Name:      input.Name,
		Email:     input.Email,
		CreatedBy: input.CreatedBy,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Cases:     nil,
		Usecases:  nil,
		Members:   nil,
	}

	if err := r.DB.Insert(&organization); err != nil {
		log.Println(err)
	}

	return &organization, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

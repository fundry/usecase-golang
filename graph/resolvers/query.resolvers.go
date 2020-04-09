package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/vickywane/usecase-server/graph/generated"
	"github.com/vickywane/usecase-server/graph/model"
)

func (r *queryResolver) User(ctx context.Context, id int) ([]*model.User, error) {
	user := model.User{ID: id}
	if err := r.DB.Select(&user); err != nil {
		return nil, err
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) (*model.User, error) {
	var users []*model.User

	if err := r.DB.Model(&users).Select(); err != nil {
		fmt.Println(err)
	}

	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Usecase(ctx context.Context, id int) (*model.Usecase, error) {
	usecase := model.Usecase{ID: id}

	if err := r.DB.Select(&usecase); err != nil {
		return nil, err
	}

	return &usecase, nil
}

func (r *queryResolver) Usecases(ctx context.Context) ([]*model.Usecase, error) {
	var usecases []*model.Usecase

	if err := r.DB.Model(&usecases).Select(); err != nil {
		log.Println(err)
	}
	return usecases, nil
}

func (r *queryResolver) Case(ctx context.Context, id int) (*model.Case, error) {
	// ====> !!Case is a reserved word in GoLang
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
	} else {
		log.Println(err, "something else")
	}

	return cases, nil
}

func (r *queryResolver) Jotter(ctx context.Context, id int) (*model.Jotter, error) {
	jotter := model.Jotter{ID: id}

	if err := r.DB.Select(&jotter); err != nil {
		return nil, err
	}

	return &jotter, nil
}

func (r *queryResolver) Jotters(ctx context.Context, usecase string) (*model.Jotter, error) {
	fmt.Println(usecase)

	jotters := model.Jotter{Usecase: usecase}
	fmt.Println(jotters)

	query := r.DB.Model(jotters).Where("Usecase == usecase")
	if err := query; err != nil {
		return nil, errors.New("error here")
	}

	return &jotters, nil
}

func (r *queryResolver) AllJotters(ctx context.Context) ([]*model.Jotter, error) {
	var jotters []*model.Jotter

	if err := r.DB.Model(&jotters).Select(); err != nil {
		return nil, err
	}

	return jotters, nil
}

func (r *queryResolver) Organization(ctx context.Context, id int) (*model.Organization, error) {
	organization := model.Organization{ID : id}

	if err := r.DB.Select(&organization); err != nil {
		return nil, err
	}

	return &organization, nil
}

func (r *queryResolver) Organizations(ctx context.Context) ([]*model.Organization, error) {
	var organizations []*model.Organization

	if err := r.DB.Model(&organizations).Select(); err != nil {
		return nil, err
	}

	return organizations, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

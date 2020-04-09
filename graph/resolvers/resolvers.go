package resolvers

import (
     "github.com/go-pg/pg/v9"

     "github.com/vickywane/usecase-server/graph/generated"
)

type Resolver struct{
     DB *pg.DB
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

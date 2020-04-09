package resolvers

import  "github.com/vickywane/usecase-server/graph/generated"

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

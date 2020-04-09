package resolvers

import (
    "context"

    "github.com/vickywane/usecase-server/graph/model"
)

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

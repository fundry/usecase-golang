package resolvers

import (
    "context"
    "fmt"
    "log"
    "math/rand"
    "time"

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

func (r *mutationResolver) UpdateOrganization(ctx context.Context,  id int, input model.UpdateOrganization) (*model.Organization, error) {
    panic(fmt.Errorf("not implemeented"))
}

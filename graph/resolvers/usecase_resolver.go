package resolvers

import (
    "context"
    "log"

    "github.com/vickywane/usecase-server/graph/model"
)

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

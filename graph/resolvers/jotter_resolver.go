package resolvers

import (
    "context"
    "errors"
    "fmt"

    "github.com/vickywane/usecase-server/graph/model"
)


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

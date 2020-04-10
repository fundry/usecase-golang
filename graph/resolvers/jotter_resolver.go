package resolvers

import (
    "context"
    "errors"
    "fmt"
    "math/rand"
    "time"

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

func (r *mutationResolver) UpdateJotter(ctx context.Context,  id int, input model.UpdateJotter) (*model.Jotter, error) {
    panic(fmt.Errorf("not implemeented"))
}

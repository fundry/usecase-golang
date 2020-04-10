package resolvers

import (
    "context"
    "fmt"
    "log"
    "math/rand"
    "time"

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

func (r *mutationResolver) UpdateUsecase(ctx context.Context,  id int,  input model.UpdateUsecase) (*model.Usecase, error) {
    panic(fmt.Errorf("not implemeented"))
}

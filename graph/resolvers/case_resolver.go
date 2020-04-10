package resolvers

import (
    "context"
    "fmt"
    "log"
    "math/rand"
    "time"

    "github.com/vickywane/usecase-server/graph/model"
)

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

func (r *mutationResolver) UpdateCase(ctx context.Context, id int, input model.UpdateCase) (*model.Case, error) {
      // cases, err :=

     panic(fmt.Errorf("not implemeented"))
}


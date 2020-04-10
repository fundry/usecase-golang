package resolvers

import (
    "context"
    "fmt"
    "math/rand"
    "time"

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

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
    user := model.User{
        ID:             rand.Int(),
        Name:           input.Name,
        Email:          input.Email,
        CreatedAt:      time.Now(),
        IsOrganization: false,
    }

    if err := r.DB.Insert(&user); err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context,  id int,  input model.UpdateUser) (*model.User, error) {
    panic(fmt.Errorf("not implemeented"))
}

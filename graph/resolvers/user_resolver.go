package resolvers

import (
    "context"
    "fmt"

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


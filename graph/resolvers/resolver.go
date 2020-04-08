package resolvers

import (
     "github.com/go-pg/pg/v9"

     "github.com/vickywane/usecase-server/graph/model"
)

type Resolver struct{
     DB *pg.DB

     users []*model.User
     cases []*model.Case
     usecases []*model.Usecase
     jotter []*model.Jotter
}

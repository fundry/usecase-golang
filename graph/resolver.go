package graph

import (
	"github.com/go-pg/pg/v9"

	"github.com/vickywane/usecase-server/graph/model"
)

type Resolver struct{
	DB *pg.DB

	users[] *model.User
	usecases[] *model.Usecase
	cases[] *model.Case
}

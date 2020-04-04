package dataloaders

import (
    "net/http"

    "github.com/go-pg/pg/v9"
)

// used in seperate dataloaders
const UserLoader = "userLoader"
const UsecaseLoader = "usecaseLoader"

func setLoader(Database *pg.DB, dataloader func(db *pg.DB, w http.ResponseWriter, r *http.Request, next http.Handler)) func(handler http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            dataloader(Database, w, r, next)
        })
    }
}

func NewMiddleware(session *pg.DB) []func(handler http.Handler) http.Handler {
    return []func(handler http.Handler) http.Handler{
        setLoader(session, User),
        setLoader(session, Usecase),
    }
}


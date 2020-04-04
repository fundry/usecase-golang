package dataloaders

import (
    "context"
    "github.com/go-pg/pg/v9"
    "net/http"
    "time"

    "github.com/vickywane/usecase-server/graph/generated"
    "github.com/vickywane/usecase-server/graph/model"
)

func User(db *pg.DB, w http.ResponseWriter, r *http.Request, next http.Handler) {
    loader := generated.NewUserLoader(generated.UserLoaderConfig{
        MaxBatch: 100,
        Wait:     1 * time.Millisecond,
        Fetch: func(keys []int) ([]*model.User, []error) {

            var dbUsers []*model.User
            err := db.Model(&dbUsers).WhereIn("id IN (?)", keys).Select()

            if err != nil {
                return []*model.User{}, []error{err}
            }

            userKeys := make(map[int]*model.User)
            users := make([]*model.User, len(keys))

            for _, user := range dbUsers {
                userKeys[11] = user
            }

            for i, k := range keys {
                if user, ok := userKeys[k]; ok {
                    users[i] = user
                }
            }

            return users, []error{err}
        },
    })

    ctx := context.WithValue(r.Context(), UserLoader, loader)
    next.ServeHTTP(w, r.WithContext(ctx))
}


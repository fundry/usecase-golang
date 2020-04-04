package dataloaders

import (
    "context"
    "github.com/go-pg/pg/v9"
    "net/http"
    "time"

    "github.com/vickywane/usecase-server/graph/generated"
    "github.com/vickywane/usecase-server/graph/model"
)

func Usecase(db *pg.DB, w http.ResponseWriter, r *http.Request, next http.Handler) {
    loader := generated.NewUserLoader(generated.UserLoaderConfig{
        MaxBatch: 100,
        Wait:     1 * time.Millisecond,
        Fetch: func(keys []int) ([]*model.Usecase, []error) {

            var dbUsers []*model.Usecase
            err := db.Model(&dbUsers).WhereIn("id IN (?)", keys).Select()

            if err != nil {
                return []*model.Usecase{}, []error{err}
            }

            // All we're doing here on out is just ordering our
            // collection to match the argument keys []int collection
            usecaseKeys := make(map[int]*model.Usecase)
            usecases := make([]*model.Usecase, len(keys))

            for _, usecase := range dbUsers {
                usecaseKeys[Usecase.ID] = usecase
            }

            for i, k := range keys {
                if usecase, ok := usecaseKeys[k]; ok {
                    usecases[i] = usecase
                }
            }

            return usecases, []error{err}
        },
    })

    ctx := context.WithValue(r.Context(), UsecaseLoader, loader)
    next.ServeHTTP(w, r.WithContext(ctx))
}


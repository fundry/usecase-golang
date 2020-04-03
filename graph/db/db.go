package db

import (
    "errors"
    "fmt"
    "github.com/go-pg/pg/v9"
    "github.com/go-pg/pg/v9/orm"
    "log"

    "github.com/vickywane/usecase-server/graph/model"
)

func Connect()  *pg.DB {
    log.Println("Db connection is starting")

   db := pg.Connect(&pg.Options{
        User:                  "postgres",
        Password:              "postgres",
        Addr:                   "localhost:5432",
        Database:              "usecase-go",
        ApplicationName:        "Usecase-server" ,
        OnConnect: func(db *pg.Conn) error {
            errors.New("unable to connect")
            return nil
        },
    })

    if db != nil {
        fmt.Println(db)
        // log.Fatalf("Unable to connect to db", db)
    }

    err := createSchema(db)
    if err != nil {
        panic(err)
    }

     return db
}

func createSchema(db *pg.DB) error {
    for _, model := range []interface{}{(*model.User)(nil) ,
        (*model.Usecase)(nil), (*model.Case)(nil)} {
        err := db.CreateTable(model, &orm.CreateTableOptions{
            IfNotExists:   true,
            Temp:          false,
            FKConstraints: true,
        })
        if err != nil {
            panic(err)
        }
    }
    return nil
}

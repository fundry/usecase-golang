package db

import (
    "fmt"
    "github.com/go-pg/pg/v9"
    "github.com/go-pg/pg/v9/orm"
    "log"

    "github.com/vickywane/usecase-server/graph/model"
)


func createSchema(db *pg.DB) error {
    for _, models := range []interface{}{(*model.User)(nil) ,
        (*model.Usecase)(nil), (*model.Case)(nil)} {
        err := db.CreateTable(models, &orm.CreateTableOptions{
            IfNotExists:true,
        })
        if err != nil {
            panic(err)
        } else {
            log.Println(err)
        }
    }
    return nil
}


func Connect()  *pg.DB {
    log.Println("Db connection is starting")

    db := pg.Connect(&pg.Options{
        User:                  "postgres",
        Password:              "postgres",
        Addr:                   "localhost:5432",
        Database:              "usecase-go",
        ApplicationName:        "Usecase-server" ,
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
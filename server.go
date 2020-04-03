package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"log"
	"net/http"
	"os"

	"github.com/vickywane/usecase-server/graph/db"
	"github.com/vickywane/usecase-server/graph"
	"github.com/vickywane/usecase-server/graph/generated"
)

const defaultPort = "4040"

//Todo connect to postgres!
//Todo Implement Auth

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var schema = handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				DB: db.Connect(),
			}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", schema)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

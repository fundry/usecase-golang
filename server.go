package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/vickywane/usecase-server/graph/dataloaders"
	"github.com/vickywane/usecase-server/graph/db"
	"github.com/vickywane/usecase-server/graph/generated"
	"github.com/vickywane/usecase-server/graph/resolvers"
)


//Todo Implement Auth

func main() {
	 defaultPort := "4040"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	Database := db.Connect()

	Database.AddQueryHook(db.Logs{})

	route := chi.NewRouter()

	route.Route("/graphql", func(route chi.Router){
		route.Use(dataloaders.NewMiddleware(Database)...)

		schema := generated.NewExecutableSchema( generated.Config{
				Resolvers: &resolvers.Resolver{
					DB: Database,
				},
				Directives:generated.DirectiveRoot{},
				Complexity:generated.ComplexityRoot{},
		 })

		var serve = handler.NewDefaultServer(schema)
		serve.Use(extension.FixedComplexityLimit(300))

		route.Handle("/", serve)
	})

	graphiql := playground.Handler("api-gateway" , "/graphql")
	route.Get("/" , graphiql)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, route))
}

package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/vickywane/usecase-server/graph/dataloaders"
	"github.com/vickywane/usecase-server/graph/db"
	"github.com/vickywane/usecase-server/graph/generated"
	InternalMiddleWare "github.com/vickywane/usecase-server/graph/middlewares"
	"github.com/vickywane/usecase-server/graph/resolvers"
)

//Todo Implement Auth --in progress
// Todo Check d & try GIN & CHI for performance

func main() {
	 defaultPort := "4040"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	Database := db.Connect()

	Database.AddQueryHook(db.Logs{})

	route := chi.NewRouter()

	route.Use(cors.New(cors.Options{
		AllowedOrigins:         []string{"http://localhost:4040"},
		AllowCredentials:       true,
		Debug:                  true,
	}).Handler)
	route.Use(middleware.Logger,
			middleware.RequestID,
			InternalMiddleWare.AuthMiddleware())

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

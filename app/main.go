package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/endot1231/ec-backend/configs"
	"github.com/endot1231/ec-backend/database"
	"github.com/endot1231/ec-backend/graph/resolvers"
	"github.com/endot1231/ec-backend/graph/services"
	"github.com/endot1231/ec-backend/internal"
	"github.com/endot1231/ec-backend/middlewares/auth"

	_ "github.com/endot1231/ec-backend/ent/runtime"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultPort = "8080"
)

func main() {
	configs.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := database.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	service := services.New(*client)
	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{
		Resolvers: &resolvers.Resolver{
			Srv: service,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	http.Handle("/query", auth.AuthMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

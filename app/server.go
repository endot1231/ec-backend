package main

import (
	"context"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/endot1231/ec-backend/graph/resolvers"
	"github.com/endot1231/ec-backend/graph/services"
	"github.com/endot1231/ec-backend/internal"
	"github.com/endot1231/ec-backend/middlewares/auth"

	"github.com/endot1231/ec-backend/ent"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const (
	defaultPort = "8080"
	dbFile      = "./db/database.db"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := ent.Open("sqlite3", "file:./database.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
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

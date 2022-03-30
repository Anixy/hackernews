package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/anixy/hackernews/graph"
	"github.com/anixy/hackernews/graph/generated"
	"github.com/anixy/hackernews/internal/auth"
	database "github.com/anixy/hackernews/internal/pkg/db/mysql"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// router := chi.NewRouter()
	// database.InitDB()
	// database.Migrate()
	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, router))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

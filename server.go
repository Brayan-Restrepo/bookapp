package main

import (
	"bookapp/config"
	"bookapp/ent"
	"bookapp/graph"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Getting().Database.Host,
		config.Getting().Database.Port,
		config.Getting().Database.User,
		config.Getting().Database.Password,
		config.Getting().Database.DBName,
		config.Getting().Database.SSLMode)

	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Client: client,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Getting().Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Getting().Server.Port, nil))
}

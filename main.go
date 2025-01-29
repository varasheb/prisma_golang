package main

import (
	"context"
	"demo/db"
	"demo/graph"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	client := db.NewClient()

	resolver := &graph.Resolver{Client: client}

	if err := client.Prisma.Connect(); err != nil {
		log.Fatalf("failed to connect to Prisma client: %v", err)
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graph.QueryType,
		Mutation: graph.MutationType,
	})
	if err != nil {
		log.Fatalf("failed to create schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphiql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "resolver", resolver)
		h.ServeHTTP(w, r.WithContext(ctx))
	}))

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	fmt.Println("Server is running on http://localhost:8000/graphiql")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatalf("failed to disconnect from Prisma client: %v", err)
		}
	}()
}

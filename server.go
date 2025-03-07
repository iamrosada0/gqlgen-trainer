package main

import (
	"gqlgen_test/generated"
	"gqlgen_test/resolver"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver.NewResolver(), // Usa a função correta para inicializar o Resolver
	}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Println("🚀 Servidor rodando em http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"log"
	"net/http"
	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/go-service-requests/graphqlconfig"
)

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig {
		Query: graphqlconfig.QueryType,
		Mutation: graphqlconfig.MutationType,
	})
	
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
	http.Handle("/", h)
	log.Println("Server running. Visit localhost:3000 for GraphQL's interface, or make HTTP requests directly.")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

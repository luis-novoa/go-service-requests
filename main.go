package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/friendsofgo/graphiql"
	"github.com/luis-novoa/go-service-requests/graphqlconfig"
)

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig {
		Query: graphqlconfig.queryType,
		Mutation: graphqlconfig.mutationType
	})

	graphQLHTTP := func(w http.ResponseWriter, r *http.Request) {
		query := graphql.Do(graphql.Params{
			Schema: schema,
			RequestString: r.URL.Query().Get("query")
		})

		if len(query.Errors) > 0 {
			fmt.Printf("errors: %v", query.Errors)
		}

		json.NewEncoder(w).Encode(query)
	}

	graphiQLInterface, err := graphiql.NewGraphiqlHandler("/")
	if err != nil { panic(err) }

	http.HandleFunc("/", graphQLHTTP)
	http.Handle("/interface", graphiQLInterface)
	log.Println("Server running. Make HTTP requests directly to localhost:3000 or access localhost:3000/interface to interact through GraphiQL.")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
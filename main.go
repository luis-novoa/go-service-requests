package main

import (
	// "fmt"
	"log"
	"net/http"
	// "encoding/json"
	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
	// "github.com/friendsofgo/graphiql"
	"github.com/luis-novoa/go-service-requests/graphqlconfig"
)

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig {
		Query: graphqlconfig.QueryType,
		Mutation: graphqlconfig.MutationType,
	})
	
	// graphiQLInterface, err := graphiql.NewGraphiqlHandler("/")
	// if err != nil { panic(err) }
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
	http.Handle("/", h)
	// http.Handle("/interface", graphiQLInterface)
	log.Println("Server running. Make HTTP requests directly to localhost:3000 or access localhost:3000/interface to interact through GraphiQL.")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// func graphQLHTTP(w http.ResponseWriter, r *http.Request) *graphql.Result {
// 	schema, _ := graphql.NewSchema(graphql.SchemaConfig {
// 		Query: graphqlconfig.QueryType,
// 		Mutation: graphqlconfig.MutationType,
// 	})

// 	query := graphql.Do(graphql.Params{
// 		Schema: schema,
// 		RequestString: r.URL.Query().Get("query"),
// 	})

// 	if len(query.Errors) > 0 {
// 		fmt.Printf("errors: %v \n", query.Errors)
// 	}

// 	return query
// 	// json.NewEncoder(w).Encode(query)
// }
package main

import (
	"log"
	"net/http"
	"github.com/graphql-go/graphql"
)

func main() {
	graphQLHTTP := func(w http.ResponseWriter, req *http.Request) {
		query := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(query)
	}

	http.HandleFunc("/", graphQLHTTP)
	log.Println("Server running on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
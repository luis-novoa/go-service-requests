package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	successMessage := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Here shall be graphQLHTTP")
	}
	http.HandleFunc("/", successMessage)
	log.Println("Server connected")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
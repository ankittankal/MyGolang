package main

import (
	"log"
	"net/http"

	"github.com/ankit/Project2-BookManagementSystem/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	/* Creates a new instance of the mux.Router from the gorilla/mux package.
	This router will be used to handle incoming HTTP requests
	and direct them to the appropriate handlers. */
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

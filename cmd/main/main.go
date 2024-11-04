package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/jitendra/pastPal/internals/routes"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	routes.RegisterAuthRoutes(r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))

}

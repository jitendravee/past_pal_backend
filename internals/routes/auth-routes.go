package routes

import (
	"github.com/gorilla/mux"
	"github.com/jitendra/pastPal/internals/controllers/auth"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/signup", auth.AddNewUserHandler).Methods("POST")
}

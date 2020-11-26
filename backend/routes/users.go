package routes

import (
	"BugTracker/handlers"

	"github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
	r := mux.NewRouter()
	// struct for user handlers
	var user handlers.UserHandlers

	// create a new user
	r.HandleFunc("/api-users/create", user.CreateUser).Methods("POST")

	return r
}

package routes

import (
	"BugTracker/handlers"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	s := r.PathPrefix("/api/users").Subrouter()
	// struct for user handlers
	var user handlers.UserHandlers

	// create a new user
	s.HandleFunc("/create/", user.CreateUser).Methods("POST")
	// login existing user
	s.HandleFunc("/login/", user.LoginUser).Methods("POST")
	// delete existing user
	s.HandleFunc("/:id/", user.DeleteUser).Methods("DELETE")
	// update user data
	s.HandleFunc("/:id/", user.UpdateUser).Methods("PUT")
}

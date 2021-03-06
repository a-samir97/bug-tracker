package routes

import (
	"BugTracker/handlers"
	"BugTracker/middlewares"

	"github.com/gorilla/mux"
)

func RolesRouter(r *mux.Router) {
	s := r.PathPrefix("/api-roles").Subrouter()

	// struct for user handlers
	var role handlers.RoleHandlers

	// create a new role
	s.HandleFunc("/create/", middlewares.TokenAuthMiddleware(role.CreateRole)).Methods("POST")

	// edit existing role
	s.HandleFunc("/edit/", middlewares.TokenAuthMiddleware(role.EditRole)).Methods("PUT")

	// delete existing role
	s.HandleFunc("/delete/", middlewares.TokenAuthMiddleware(role.EditRole)).Methods("DELETE")

}

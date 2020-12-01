package routes

import (
	"BugTracker/handlers"

	"github.com/gorilla/mux"
)

func RolesRouter(r *mux.Router) {
	s := r.PathPrefix("/api-roles").Subrouter()

	// struct for user handlers
	var role handlers.RoleHandlers

	// create a new role
	s.HandleFunc("/create/", role.CreateRole).Methods("POST")

	// edit existing role
	s.HandleFunc("/edit/", role.EditRole).Methods("PUT")

	// delete existing role
	s.HandleFunc("/delete/", role.DeleteRole).Methods("DELETE")

}

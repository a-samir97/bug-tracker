package routes

import (
	"BugTracker/handlers"

	"github.com/gorilla/mux"
)

func RolesRouter() *mux.Router {
	r := mux.NewRouter()

	// struct for user handlers
	var role handlers.RoleHandlers

	// create a new role
	r.HandleFunc("/api-roles/create/", role.CreateRole).Methods("POST")
	// edit existing role
	r.HandleFunc("/api-roles/edit", role.EditRole).Methods("PUT")
	// delete existing role
	r.HandleFunc("/api-roles/delete", role.DeleteRole).Methods("DELETE")
	return r
}

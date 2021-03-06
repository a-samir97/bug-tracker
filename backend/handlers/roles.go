package handlers

import (
	"BugTracker/models"
	"encoding/json"
	"net/http"
)

type RoleHandlers struct{}

// CreateRole ... for creating role (admin, stuff, customer)
func (role *RoleHandlers) CreateRole(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()

	if err != nil {
		w.Write([]byte("Can not decode JSON data"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newRole, created, err := models.RoleDb.GetOrInsertRole(data["role_name"])

	if err != nil {
		w.Write([]byte("Can not encode to JSON"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(&newRole)

	if err != nil {
		w.Write([]byte("Can not encode to JSON"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if created {
		w.Write([]byte(jsonData))
		w.WriteHeader(http.StatusCreated)
	} else {
		w.Write([]byte(jsonData))
		w.WriteHeader(http.StatusOK)
	}

}

func (role *RoleHandlers) DeleteRole(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (role *RoleHandlers) EditRole(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

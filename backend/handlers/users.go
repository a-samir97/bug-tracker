package handlers

import (
	"BugTracker/auth"
	"BugTracker/helpers"
	"BugTracker/models"
	"BugTracker/models/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlers struct{}

// CreateUser function ... to create a new user
func (u *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {

	// User Data
	// username, email, password, user_role
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()

	password := data["password"]

	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// check if username is not exists
	usernameExists := helpers.UsernameExists(data["username"])
	if usernameExists {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": "username is already exist"})
		return
	}

	// check if email is not exists
	emailExists := helpers.EmailExists(data["email"])
	if emailExists {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": "email is already exist"})
		return
	}

	// encrypt user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user := &models.User{
		Username:  data["username"],
		Password:  string(hashedPassword),
		Email:     data["email"],
		Role:      data["role"],
		FirstName: data["first_name"],
		LastName:  data["last_name"]}

	_, err = sql.UserDb.InsertUser(user)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
		return
	}

}

//LoginUser function .. to login user usering jwt authentication
func (u *UserHandlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	// User Data
	// email, password
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()

	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	user := helpers.AuthenticateUser(data["email"], data["password"])

	if user != nil {
		token, err := auth.CreateToken(user.Id)

		if err != nil {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		} else {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"token": token})
		}
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	// TODO: add constant error messages
	json.NewEncoder(w).Encode(map[string]string{"detail": "make sure that your credentials are right"})
}

// TODO: update user information
func (u *UserHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// TODO: get user information by ID
func (u *UserHandlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get user id from URL param
	vars := mux.Vars(r)
	userId := vars["id"]
	pasredUserId, err := strconv.Atoi(userId)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"detail": "Please make sure that put a valid User ID."})
		return
	}

	// need to change
	fetchedUser, err := sql.UserDb.GetUser(pasredUserId)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fetchedUser)
}

// TODO: Delete user by ID
func (u *UserHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get User Id
	vars := mux.Vars(r)
	userId := vars["id"]
	pasredUserId, err := strconv.Atoi(userId)

	// if there is error in parsing
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"detail": err.Error()})
	}

	err = sql.UserDb.DeleteUser(pasredUserId)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"detail": err.Error()})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

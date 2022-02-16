package handlers

import (
	"BugTracker/auth"
	"BugTracker/helpers"
	"BugTracker/models"
	"encoding/json"
	"net/http"

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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is something wrong, please try again"))
	}

	// check if username is not exists
	usernameExists := helpers.UsernameExists(data["username"])
	if usernameExists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username is already exist"))
		return
	}

	// check if email is not exists
	emailExists := helpers.EmailExists(data["email"])
	if emailExists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("email is already exist"))
		return
	}

	// encrypt user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is something wrong, please try again"))
		return
	}

	user := &models.User{
		Username:  data["username"],
		Password:  string(hashedPassword),
		Email:     data["email"],
		Role:      data["role"],
		FirstName: data["first_name"],
		LastName:  data["last_name"]}

	_, err = models.UserDb.InsertUser(user)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is something wrong, please try again"))
	}

	user := helpers.AuthenticateUser(data["email"], data["password"])

	if user != nil {
		token, err := auth.CreateToken(user.Id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("there is something wrong, please try again"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(token))
		}
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("make sure that your credentials are right"))
}

func (u *UserHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandlers) GetUserByID(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

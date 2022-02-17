package test

import (
	"BugTracker/handlers"
	"BugTracker/helpers"
	"BugTracker/models"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: move this methods to helpers
func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

var (
	createUserURL = "/api/users/create"
	loginURL      = "/api/users/login/"
	userHandlers  = handlers.UserHandlers{}
)

func TestCreateUserSuccessCase(t *testing.T) {
	db, _ := helpers.InitSQLiteDB()
	models.UserDb.Db = db
	data := []byte(`{"username":"ahmedsamir", "email":"ahmedd@gmail.com", "password":"a.samir1199", "first_name":"Ahmed", "last_name":"Samir", "role":"Admin"}`)
	request, err := http.NewRequest("POST", createUserURL, bytes.NewBuffer(data))

	if err != nil {
		log.Println(err.Error())
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandlers.CreateUser)
	handler.ServeHTTP(response, request)
	assertEqual(t, http.StatusCreated, response.Code, "")
	defer db.Close()
	// TODO: remove all database after unittests
}

func TestCreateUserFailureCase(t *testing.T) {
	// not valid data, don't pass email and password
	// should give 400 bad request
	db, _ := helpers.InitSQLiteDB()
	models.UserDb.Db = db

	data := []byte(`"username":"ahmedsamir"`)
	request, _ := http.NewRequest("POST", createUserURL, bytes.NewBuffer(data))
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandlers.CreateUser)
	handler.ServeHTTP(response, request)
	fmt.Println(response.Code)
	assertEqual(t, http.StatusBadRequest, response.Code, "")

	defer db.Close()

	// TODO: remove all database after unittests
}

// func TestLoginSuccessCase(t *testing.T) {

// }

// func TestLoginFailureCase(t *testing.T) {

// }

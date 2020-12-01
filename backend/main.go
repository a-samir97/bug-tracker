package main

import (
	"BugTracker/helpers"
	"BugTracker/models"
	"BugTracker/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	DB, err := helpers.InitDb()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	models.RoleDb.Db = DB
	models.BugDb.Db = DB
	models.UserDb.Db = DB
	models.StatusDb.Db = DB

	r := mux.NewRouter()

	routes.RolesRouter(r)
	routes.UserRouter(r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

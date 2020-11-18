package main

import (
	"BugTracker/helpers"
	"BugTracker/models"
	"BugTracker/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := helpers.InitDb()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	role := models.RoleDB{Db: db}
	_ = role.GetOrInsertRole("admin")

	r := routes.Router()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

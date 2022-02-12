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

	// for developing and testing use SQLite3
	// for production use PostgreSQL
	DB, err := helpers.InitSQLiteDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	// need better implementation for this part
	// reed about design pattern, to solve this part
	models.BugDb.Db = DB
	models.UserDb.Db = DB

	r := mux.NewRouter()

	routes.UserRouter(r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server is running or port 8000 ...")
	log.Fatal(srv.ListenAndServe())
}

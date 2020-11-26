package helpers

import (
	"BugTracker/config"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // sql behavior modified
)

func InitDb() (*sql.DB, error) {
	var connectionString string

	if os.Getenv("GO111MODULE") == "travis" {
		connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			"127.0.0.1", 5432, "travis", "", "travis_ci_test")
	} else {
		connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Dbname)
	}

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return db, nil
}

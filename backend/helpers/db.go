package helpers

import (
	"BugTracker/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // sql behavior modified
)

func InitDb() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return db, nil
}

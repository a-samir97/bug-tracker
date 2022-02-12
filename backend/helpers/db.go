package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // sql behavior modified (postgreSQL database)
	_ "github.com/mattn/go-sqlite3" // for sqlite3 database
)

func InitDb() (*sql.DB, error) {
	var connectionString string
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return db, nil
}

// to initize SQLite database for developing and testing
func InitSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "bug_project.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}

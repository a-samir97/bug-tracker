package sql

import (
	"database/sql"
)

type BugDB struct {
	Db *sql.DB
}

var BugDb = &BugDB{}

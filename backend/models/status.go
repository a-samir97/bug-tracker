package models

import "database/sql"

// Status struct ...
type Status struct {
	ID   int    `json:"status_id"`
	Name string `json:"status_name"`
}

type StatusDB struct {
	Db *sql.DB
}

var StatusDb = &StatusDB{}

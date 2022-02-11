package models

import (
	"database/sql"
)

// Bug struct ...
type Bug struct {
	Base
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Screenshot  byte      `json:"screenshot"`
	User        User      `json:"user"`
	Status      string    `json:"status"`
}

func (bug *Bug) getAssignedUsers() {

}

func (bug *Bug) getBugStatus() {

}

func (bug *Bug) getCreatedUser() {

}

type BugDB struct {
	Db *sql.DB
}

var BugDb = &BugDB{}

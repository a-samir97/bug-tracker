package models

import (
	"database/sql"
	"time"
)

// Bug struct ...
type Bug struct {
	ID          int       `json:"bug_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Screenshot  byte      `json:"screenshot"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	Status      Status    `json:"status"`
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

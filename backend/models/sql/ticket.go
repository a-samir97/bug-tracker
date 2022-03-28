package sql

import (
	"database/sql"
)

// need to do more cleanup here
type TicketDB struct {
	Db *sql.DB
}

var TicketDb = &TicketDB{}

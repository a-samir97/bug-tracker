package models

import "database/sql"

// Role struct ...
// Admin or Stuff or Customer
type Role struct {
	ID   int    `json:"role_id"`
	Name string `json:"role_name"`
}

// RoleDB to handle queries in database
type RoleDB struct {
	db *sql.DB
}

func (role *RoleDB) insertNewRole() {

}

func (role *RoleDB) deleteRole() {

}

func (role *RoleDB) updateRole() {

}

package models

import (
	"database/sql"
)

// Role struct ...
// Admin or Stuff or Customer
type Role struct {
	ID   int    `json:"role_id"`
	Name string `json:"role_name"`
}

// RoleDB to handle queries in database
type RoleDB struct {
	Db *sql.DB
}

// GetOrInsertRole ...
// to insert a new role or get existing role
func (role *RoleDB) GetOrInsertRole(userRole string) (*Role, bool, error) {
	var roleID int
	var roleName string

	_ = role.Db.QueryRow("SELECT role_id, role_name from roles WHERE role_name = $1 ", userRole).Scan(&roleID, &roleName)

	// insert a new role and return id, role name after inserting it in the database
	if roleID == 0 && roleName == "" {
		sqlStatement, err := role.Db.Prepare(`INSERT INTO roles (role_name) VALUES ($1) RETURNING role_id, role_name`)
		if err != nil {
			return nil, false, err
		}

		// use Exec
		_ = sqlStatement.QueryRow(userRole).Scan(&roleID, &roleName)
		roleStruct := &Role{ID: roleID, Name: roleName}
		return roleStruct, true, nil
	}

	// return existing role
	roleStruct := &Role{ID: roleID, Name: roleName}
	return roleStruct, false, nil
}

func (role *RoleDB) GetRoleByID(roleId string) *Role {
	var roleID int
	var roleName string

	_ = role.Db.QueryRow("SELECT role_id, role_name from roles WHERE role_id = $1", roleId).Scan(&roleID, &roleName)

	roleStruct := &Role{ID: roleID, Name: roleName}

	return roleStruct
}

func (role *RoleDB) deleteRole() {

}

// RoleDb.. database instance for roles
var RoleDb = &RoleDB{}

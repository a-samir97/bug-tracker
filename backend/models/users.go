package models

import (
	"database/sql"
	"time"
)

// User struct ....
type User struct {
	ID        int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Role      Role      `json:"role"`
}

func (user *User) getUsername() string {
	return user.Username
}

func (user *User) getEmail() string {
	return user.Email
}

func (user *User) getUserRole() string {
	return user.Role.Name
}

func (user *User) isAdmin() bool {
	if user.Role.Name == "Admin" {
		return true
	}
	return false
}

func (user *User) isStuff() bool {
	if user.Role.Name == "Stuff" {
		return true
	}
	return false
}

func (user *User) isCustomer() bool {
	if user.Role.Name == "Customer" {
		return true
	}
	return false
}

type UserDB struct{}

func (user *UserDB) insertUser(db *sql.DB, newUser User) error {
	return nil
}

func (user *UserDB) updateUser(db *sql.DB, updatedUser User) *User {

	return nil
}

func (user *UserDB) getUser(db *sql.DB, id int) *User {

	return nil
}

func (user *UserDB) deleteUser(db *sql.DB, id int) error {

	return nil
}

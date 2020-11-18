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

type UserDB struct {
	Db *sql.DB
}

func (user *UserDB) InsertUser(newUser *User) (*User, error) {

	return nil, nil
}

func (user *UserDB) UpdateUser(updatedUser User) *User {

	return nil
}

func (user *UserDB) GetUser(id int) *User {

	return nil
}

func (user *UserDB) DeleteUser(id int) error {

	return nil
}

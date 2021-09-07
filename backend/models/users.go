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
	Role      Role      `json:"role_id"`
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
	sqlStatement, err := UserDb.Db.Prepare(`
	INSERT INTO users (username, password, email, role_id, created_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING username, email, role_id `)

	if err != nil {
		return nil, err
	}
	_, err = sqlStatement.Exec(newUser.Username, newUser.Password, newUser.Email, newUser.Role.ID, newUser.CreatedAt)

	if err != nil {
		return nil, err
	}

	return newUser, nil
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

func (user *UserDB) ChangePassword() error {

	return nil
}

var UserDb = &UserDB{}

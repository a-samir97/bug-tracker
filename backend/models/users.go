package models

import (
	"database/sql"
)

// User struct ....
type User struct {
	Base
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

func (user *User) getUsername() string {
	return user.Username
}

func (user *User) getEmail() string {
	return user.Email
}

func (user *User) GetUserRole() string {
	return user.Role
}

// TOOD: Re-do roles of users
func (user *User) IsAdmin() bool {
	return user.Role == "Admin"
}

func (user *User) IsStuff() bool {
	return user.Role == "Stuff"
}

type UserDB struct {
	Db *sql.DB
}

// This method responsible for creating a new user
func (user *UserDB) InsertUser(newUser *User) (*User, error) {
	sqlStatement, err := UserDb.Db.Prepare(`
	INSERT INTO users (first_name, last_name, username, password, email, role) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING username, email, role `)

	if err != nil {
		return nil, err
	}
	_, err = sqlStatement.Exec(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Password, newUser.Email, newUser.Role)

	if err != nil {
		return nil, err
	}
	defer sqlStatement.Close()
	return newUser, nil
}

// This method responsible for update user info
func (user *UserDB) UpdateUser(updatedUser User) error {
	stmt, err := UserDb.Db.Prepare(`UPDATE users set username = ?, email = ? WHERE id = ?`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(updatedUser.getUsername(), updatedUser.getEmail(), updatedUser.Id)

	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

// This method responsible for deleting a user
func (user *UserDB) GetUser(id int) error {
	stmt, err := UserDb.Db.Prepare(`SELECT username, first_name, last_name, email FROM users WHERE id = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil
}

func (user *UserDB) DeleteUser(id int) error {

	stmt, err := UserDb.Db.Prepare(`DELETE FROM users WHERE id = ?`)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil
}

func (user *UserDB) ChangePassword() error {

	return nil
}

var UserDb = &UserDB{}

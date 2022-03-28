package sql

import (
	"BugTracker/models"
	"database/sql"
)

type UserDB struct {
	Db *sql.DB
}

var UserDb = &UserDB{}

// This method responsible for creating a new user
func (user *UserDB) InsertUser(newUser *models.User) (*models.User, error) {
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
func (user *UserDB) UpdateUser(updatedUser models.User) error {
	stmt, err := UserDb.Db.Prepare(`UPDATE users set username = ?, email = ? WHERE id = ?`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(updatedUser.GetUsername(), updatedUser.GetEmail(), updatedUser.Id)

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

// TODO: Change password for existing user
func (user *UserDB) ChangePassword() error {

	return nil
}

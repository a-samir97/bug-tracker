package helpers

import (
	"BugTracker/models"
	"BugTracker/models/sql"

	"golang.org/x/crypto/bcrypt"
)

// UsernameExists ... to check for username if exists or not ...
func UsernameExists(username string) bool {
	var usernameUser string

	_ = sql.UserDb.Db.QueryRow("SELECT username FROM users WHERE username = $1", username).Scan(&usernameUser)

	return (usernameUser != "")
}

// EmailExists ...to check for email if exists or not ...
func EmailExists(email string) bool {
	var checkEmail string

	_ = sql.UserDb.Db.QueryRow("SELECT email FROM users WHERE email = $1", email).Scan(&checkEmail)

	return (checkEmail != "")
}

// HashPassword function .. to encrypt user's password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// CheckPassword function ... to check if the password is correct
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// AuthenticateUser ...
func AuthenticateUser(email string, password string) *models.User {
	var userEmail string
	var hashPassword string
	var userID int

	_ = sql.UserDb.Db.QueryRow("SELECT user_id, email, password FROM users WHERE email = $1", email).Scan(&userID, &userEmail, &hashPassword)
	if userEmail == "" || hashPassword == "" {
		return nil
	}

	checkUserPassword := CheckPassword(password, hashPassword)

	if checkUserPassword {
		user := models.User{Email: userEmail}
		return &user
	}

	return nil
}

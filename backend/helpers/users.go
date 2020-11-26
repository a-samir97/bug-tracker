package helpers

import "BugTracker/models"

// to check for username if exists or not ...
func UsernameExists(username string) bool {
	var usernameUser string

	_ = models.UserDb.Db.QueryRow("SELECT username FROM users WHERE username = $1", username).Scan(&usernameUser)

	if usernameUser == "" {
		return false
	}
	return true
}

// to check for email if exists or not ...
func EmailExists(email string) bool {
	var checkEmail string

	_ = models.UserDb.Db.QueryRow("SELECT email FROM users WHERE email = $1", email).Scan(&checkEmail)

	if checkEmail == "" {
		return false
	}
	return true
}

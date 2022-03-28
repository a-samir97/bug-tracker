package models

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

func (user *User) GetUsername() string {
	return user.Username
}

func (user *User) GetEmail() string {
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

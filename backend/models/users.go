package models

import "time"

// User struct ....
type User struct {
	ID        int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Role      Role      `json:"role"`
}

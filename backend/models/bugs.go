package models

import (
	"time"
)

// Bug struct ...
type Bug struct {
	ID          int       `json:"bug_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Screenshot  byte      `json:"screenshot"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	Status      Status    `json:"status"`
}

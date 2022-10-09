package model

import "time"

// User information
type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

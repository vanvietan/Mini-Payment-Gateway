package model

import (
	"gorm.io/gorm"
	"time"
)

// Transaction handle transaction and otp
type Transaction struct {
	ID        int64
	CardID    int64
	OrderID   int64
	OTP       string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

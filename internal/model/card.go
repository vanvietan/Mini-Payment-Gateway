package model

import (
	"gorm.io/gorm"
	"time"
)

// Card user card use for payment
type Card struct {
	ID          int64
	Number      string
	ExpiredDate time.Time
	CVV         string
	Balance     int64
	UserID      int64
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

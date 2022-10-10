package model

import (
	"gorm.io/gorm"
	"time"
)

// Card user card use for payment
type Card struct {
	ID          int64
	ExpiredDate time.Time
	CVV         int16
	Balance     int64
	UserID      int64
	DeletedAt   gorm.DeletedAt
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AuditTrails []AuditTrail `gorm:"foreignKey:CardID;references:ID"`
}

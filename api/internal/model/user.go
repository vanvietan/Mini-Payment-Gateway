package model

import "time"

// User information
type User struct {
	ID          int64
	Username    string
	Password    string
	Cards       []Card `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AuditTrails []AuditTrail `gorm:"foreignKey:UserID;references:ID"`
}

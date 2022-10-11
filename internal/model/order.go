package model

import "time"

// Order model
type Order struct {
	ID          int64
	Amount      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AuditTrails []AuditTrail `gorm:"foreignKey:OrderID;references:ID"`
}

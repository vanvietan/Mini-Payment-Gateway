package model

import "time"

// AuditTrail tracking card status by admin depends on
type AuditTrail struct {
	ID        int64
	CardID    int64
	UserID    int64
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

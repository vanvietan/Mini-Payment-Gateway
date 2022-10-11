package model

type Status string

/*
Enums for Status in audit_trail table
*/

const (
	StatusAccepted Status = "ACCEPTED"
	StatusPending  Status = "PENDING"
	StatusDeleted  Status = "DELETED"
)

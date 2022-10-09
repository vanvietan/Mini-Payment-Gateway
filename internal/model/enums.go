package model

type Status string

/*
Enums for Status in audit_trail table
*/

const (
	StatusInUse   Status = "IN_USE"
	StatusExpired Status = "EXPIRED"
	StatusDeleted Status = "DELETED"
)

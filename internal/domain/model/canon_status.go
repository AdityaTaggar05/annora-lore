package model

type CanonStatus string

const (
	CanonStatusAccepted CanonStatus = "accepted"
	CanonStatusRejected CanonStatus = "rejected"
	CanonStatusPending  CanonStatus = "pending"
)

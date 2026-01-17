package model

import "time"

type LoreNode struct {
	ID          string
	Type        NodeType
	Name        string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CanonStatus CanonStatus
	WorldID     string
	Custom      map[string]any
}

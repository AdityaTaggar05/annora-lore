package model

import "time"

type LoreNode struct {
	ID              string      `json:"id"`
	Type            NodeType    `json:"type"`
	Name            string      `json:"name"`
	CreatedBy       string      `json:"created_by"`
	CreatorUsername string      `json:"creator_username"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	CanonStatus     CanonStatus `json:"canon_status"`
	WorldID         string      `json:"world_id"`
}

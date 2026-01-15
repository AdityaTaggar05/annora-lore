package models

type LoreNode struct {
	ID          string         `json:"id" validate:"required,uuid"`
	Type        NodeType       `json:"type" validate:"required, oneof=place species character artifact faction event"`
	Name        string         `json:"name" validate:"required,min=1,max=50"`
	CreatedBy   string         `json:"created_by" validate:"required,uuid"`
	CanonStatus CanonStatus    `json:"canon_status" validate:"required, oneof=accepted rejected pending"`
	WorldID     string         `json:"world_id" validate:"required,uuid"`
	Links       []Link         `json:"links,omitempty" validate:"-"`
	Custom      map[string]any `json:"custom,omitempty" validate:"-"`
}

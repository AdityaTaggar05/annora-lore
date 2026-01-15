package models

type Link struct {
	TargetID string `json:"target_id" validate:"required,uuid"`
	Relation string `json:"relation" validate:"required,min=1,max=50"`
}

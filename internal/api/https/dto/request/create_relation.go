package request

import (
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

type CreateRelationRequestDTO struct {
	WorldID     string `json:"world_id" validate:"required,uuid"`
	CreatedBy   string `json:"created_by" validate:"required,uuid"`
	FromID      string `json:"from_id" validate:"required,uuid"`
	ToID        string `json:"to_id" validate:"required,uuid"`
	FromType    string `json:"from_type" validate:"required,oneof=place species character faction event artifact"`
	ToType      string `json:"to_type" validate:"required,oneof=place species character faction event artifact"`
	Relation    string `json:"relation" validate:"required"`
	Description string `json:"description" validate:"required,min=2,max=100"`
}

func (dto *CreateRelationRequestDTO) ToDomain() (string, model.Relation) {
	now := time.Now()

	return dto.WorldID, model.Relation{
		FromID:      dto.FromID,
		ToID:        dto.ToID,
		Rel:         model.RelationType(dto.Relation),
		Description: dto.Description,
		CreatedBy:   dto.CreatedBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

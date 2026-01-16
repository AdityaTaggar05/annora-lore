package request

import (
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/google/uuid"
)

type CreateLoreNodeRequestDTO struct {
	Name      string         `json:"name" validate:"required,min=2,max=100"`
	Type      model.NodeType `json:"type" validate:"required,oneof=place species character artifact faction event"`
	WorldID   string         `json:"world_id" validate:"required,uuid"`
	CreatedBy string         `json:"created_by" validate:"required,uuid"`
}

func (dto CreateLoreNodeRequestDTO) ToDomain() model.LoreNode {
	now := time.Now()

	return model.LoreNode{
		ID:          uuid.NewString(),
		Type:        dto.Type,
		Name:        dto.Name,
		CreatedBy:   dto.CreatedBy,
		CreatedAt:   now,
		UpdatedAt:   now,
		CanonStatus: model.CanonStatusPending,
		WorldID:     dto.WorldID,
	}
}

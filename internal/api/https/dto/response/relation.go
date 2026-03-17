package response

import (
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

type LoreRelationResponseDTO struct {
	FromID      string    `json:"from_id"`
	ToID        string    `json:"to_id"`
	Rel         string    `json:"rel"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToLoreRelationResponse(rel model.Relation) LoreRelationResponseDTO {
	return LoreRelationResponseDTO{
		FromID:      rel.FromID,
		ToID:        rel.ToID,
		Rel:         string(rel.Rel),
		Description: rel.Description,
		CreatedBy:   rel.CreatedBy,
		CreatedAt:   rel.CreatedAt,
		UpdatedAt:   rel.UpdatedAt,
	}
}

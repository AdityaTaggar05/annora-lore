package response

import (
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

type LoreRelationResponseDTO struct {
	ToID            string    `json:"to_id"`
	Rel             string    `json:"rel"`
	Description     string    `json:"description"`
	CreatorUsername string    `json:"creator_username"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func ToLoreRelationResponse(rel model.Relation) LoreRelationResponseDTO {
	return LoreRelationResponseDTO{
		ToID:            rel.ToID,
		Rel:             string(rel.Rel),
		Description:     rel.Description,
		CreatorUsername: rel.CreatorUsername,
		CreatedAt:       rel.CreatedAt,
		UpdatedAt:       rel.UpdatedAt,
	}
}

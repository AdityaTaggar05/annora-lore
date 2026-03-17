package response

import (
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

type LoreNodeResponseDTO struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	Name        string         `json:"name"`
	CreatedBy   string         `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	CanonStatus string         `json:"canon_status"`
	WorldID     string         `json:"world_id"`
}

func ToLoreNodeResponse(node model.LoreNode) LoreNodeResponseDTO {
	return LoreNodeResponseDTO{
		ID:          node.ID,
		Type:        string(node.Type),
		Name:        node.Name,
		CreatedBy:   node.CreatedBy,
		CreatedAt:   node.CreatedAt,
		CanonStatus: string(node.CanonStatus),
		WorldID:     node.WorldID,
	}
}

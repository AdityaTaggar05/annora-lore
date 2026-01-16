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
	UpdatedAt   time.Time      `json:"updated_at"`
	CanonStatus string         `json:"canon_status"`
	WorldID     string         `json:"world_id"`
	Custom      map[string]any `json:"custom,omitempty"`
}

func ToLoreNodeResponse(node model.LoreNode) LoreNodeResponseDTO {
	return LoreNodeResponseDTO{
		ID:          node.ID,
		Type:        string(node.Type),
		Name:        node.Name,
		CreatedBy:   node.CreatedBy,
		CreatedAt:   node.CreatedAt,
		UpdatedAt:   node.UpdatedAt,
		CanonStatus: string(node.CanonStatus),
		WorldID:     node.WorldID,
		Custom:      node.Custom,
	}
}

package response

import (
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

type LoreNodeResponseDTO struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	Name            string    `json:"name"`
	CreatorUsername string    `json:"creator_username"`
	CreatedAt       time.Time `json:"created_at"`
	CanonStatus     string    `json:"canon_status"`
}

func ToLoreNodeResponse(node model.LoreNode) LoreNodeResponseDTO {
	return LoreNodeResponseDTO{
		ID:              node.ID,
		Type:            string(node.Type),
		Name:            node.Name,
		CreatorUsername: node.CreatorUsername,
		CreatedAt:       node.CreatedAt,
		CanonStatus:     string(node.CanonStatus),
	}
}

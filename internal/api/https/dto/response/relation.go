package response

import "github.com/AdityaTaggar05/annora-lore/internal/domain/model"

type LoreRelationResponseDTO struct {
	From     model.LoreNode `json:"from"`
	Relation model.Relation `json:"relation"`
	To       model.LoreNode `json:"to"`
}

func ToLoreRelationResponse(from model.LoreNode, rel model.Relation, to model.LoreNode) LoreRelationResponseDTO {
	return LoreRelationResponseDTO{
		From:     from,
		Relation: rel,
		To:       to,
	}
}

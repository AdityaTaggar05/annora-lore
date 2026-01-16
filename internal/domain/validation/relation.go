package validation

import (
	"errors"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

func ValidateRelation(rel model.Relation) error {
	if rel.Rel == "RELATED_TO" {
		if len(rel.Description) < 10 {
			return errors.New("RELATED_TO requires a meaningful description")
		}
		return nil
	}

	toMap, ok := model.AllowedRelations[rel.From]
	if !ok {
		return errors.New("no relations allowed from node type: " + string(rel.From))
	}

	relMap, ok := toMap[rel.To]
	if !ok || !relMap[rel.Rel] {
		return errors.New("invalid relation to node type: " + string(rel.To) + " from node type: " + string(rel.From))
	}

	return nil
}

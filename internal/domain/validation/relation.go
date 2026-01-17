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

	toMap, ok := model.AllowedRelations[rel.FromType]
	if !ok {
		return errors.New("no relations allowed from node type: " + string(rel.FromType))
	}

	relMap, ok := toMap[rel.ToType]
	if !ok || !relMap[rel.Rel] {
		return errors.New("invalid relation to node type: " + string(rel.ToType) + " from node type: " + string(rel.FromType))
	}

	return nil
}

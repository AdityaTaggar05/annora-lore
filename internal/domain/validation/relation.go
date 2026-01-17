package validation

import (
	"errors"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

func ValidateRelation(fromType, toType model.NodeType, rel model.RelationType) error {
	toMap, ok := model.AllowedRelations[fromType]
	if !ok {
		return errors.New("no relations allowed from node type: " + string(fromType))
	}

	relMap, ok := toMap[toType]
	if !ok || !relMap[rel] {
		return errors.New("invalid relation to node type: " + string(toType) + " from node type: " + string(fromType))
	}

	return nil
}

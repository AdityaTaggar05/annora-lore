package validation

import (
	"errors"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

func ValidateNodeType(t string) error {
	if model.AllowedTypes[model.NodeType(t)] {
		return nil
	}

	return errors.New("no such node type exists: " + t)
}

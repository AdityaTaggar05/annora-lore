package service

import (
	"context"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/validation"
)

func (s *LoreService) FetchNodeByID(ctx context.Context, worldID, nodeID string) (*model.LoreNode, error) {
	return s.Repo.FetchNodeByID(ctx, worldID, nodeID)
}

func (s *LoreService) FetchNodesByType(ctx context.Context, worldID string, nodeType []string) ([]*model.LoreNode, error) {
	for _, t := range nodeType {
		if err := validation.ValidateNodeType(t); err != nil {
			return nil, err
		}
	}

	return s.Repo.FetchNodesByType(ctx, worldID, nodeType)
}

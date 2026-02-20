package service

import (
	"context"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

func (s *LoreService) FetchNodeByID(ctx context.Context, worldID, nodeID string) (*model.LoreNode, error) {
	return s.Repo.FetchNodeByID(ctx, worldID, nodeID)
}

package service

import (
	"context"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

func (s *LoreService) FetchAllRelationsOfNode(ctx context.Context, nodeID string) ([]*model.Relation, error) {
	return s.Repo.FetchAllRelationsOfNode(ctx, nodeID)
}

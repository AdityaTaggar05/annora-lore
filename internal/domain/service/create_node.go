package service

import (
	"context"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/google/uuid"
)

func (s *LoreService) CreateNode(ctx context.Context) error {
	now := time.Now()

	node := model.LoreNode{
		ID:          uuid.New().String(),
		Type:        "place",
		Name:        uuid.New().String(),
		WorldID:     "world1",
		CreatedBy:   uuid.New().String(),
		CanonStatus: model.CanonStatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return s.Repo.CreateNode(ctx, node)
}

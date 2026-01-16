package services

import (
	"context"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/models"
	"github.com/google/uuid"
)

func (s *LoreService) CreateNode(ctx context.Context) error {
	now := time.Now()

	node := models.LoreNode{
		ID:          uuid.New().String(),
		Type:        "place",
		Name:        uuid.New().String(),
		WorldID:     "world1",
		CreatedBy:   uuid.New().String(),
		CanonStatus: models.CanonStatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return s.Repo.CreateNode(ctx, node)
}

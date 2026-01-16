package service

import (
	"context"

	"github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/request"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

func (s *LoreService) CreateNode(ctx context.Context, dto request.CreateLoreNodeRequestDTO) (*model.LoreNode, error) {
	return s.Repo.CreateNode(ctx, dto.ToDomain())
}

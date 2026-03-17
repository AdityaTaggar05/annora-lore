package service

import (
	"context"

	"github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/request"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
)

func (s *LoreService) CreateNode(ctx context.Context, dto request.CreateLoreNodeRequestDTO) (*model.LoreNode, error) {
	// TODO: Fetch Username from User Service corresponding to user_id from JWT Token
	// TODO: Use Redis to cache the usernames from User Service

	return s.Repo.CreateNode(ctx, dto.ToDomain("username"))
}

package service

import (
	"context"
	"errors"

	"github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/request"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/validation"
)

func (s *LoreService) CreateRelation(ctx context.Context, dto request.CreateRelationRequestDTO) (*model.Relation, error) {
	// TODO: Fetch Username from User Service corresponding to user_id from JWT Token
	// TODO: Use Redis to cache the usernames from User Service

	worldID, relation := dto.ToDomain("username")

	if dto.FromID == dto.ToID {
		return nil, errors.New("cannot create a self-pointing relation")
	}

	err := validation.ValidateRelation(model.NodeType(dto.FromType), model.NodeType(dto.ToType), model.RelationType(dto.Relation))
	if err != nil {
		return nil, err
	}

	return s.Repo.CreateRelation(ctx, worldID, &relation)
}

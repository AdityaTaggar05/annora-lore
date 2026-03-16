package service

import (
	"context"
	"errors"

	"github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/request"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/validation"
)

func (s *LoreService) CreateRelation(ctx context.Context, dto request.CreateRelationRequestDTO) (*model.LoreNode, *model.Relation, *model.LoreNode, error) {
	worldID, relation := dto.ToDomain()

	if dto.FromID == dto.ToID {
		return nil, nil, nil, errors.New("cannot create a self-pointing relation")
	}

	err := validation.ValidateRelation(model.NodeType(dto.FromType), model.NodeType(dto.ToType), model.RelationType(dto.Relation))
	if err != nil {
		return nil, nil, nil, err
	}

	return s.Repo.CreateRelation(ctx, worldID, relation)
}

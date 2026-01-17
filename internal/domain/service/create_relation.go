package service

import (
	"context"

	"github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/request"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/validation"
)

func (s *LoreService) CreateRelation(ctx context.Context, dto request.CreateRelationRequestDTO) (*model.LoreNode, *model.Relation, *model.LoreNode, error) {
	worldID, relation := dto.ToDomain()

	err := validation.ValidateRelation(model.NodeType(dto.FromType), model.NodeType(dto.ToType), model.RelationType(dto.Relation))
	if err != nil {
		return nil, nil, nil, err
	}

	return s.Repo.CreateRelation(ctx, worldID, relation)
}

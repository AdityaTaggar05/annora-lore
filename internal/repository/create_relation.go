package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (r *LoreRepository) CreateRelation(ctx context.Context, worldID string, relation *model.Relation) (*model.Relation, error) {
	_, err := r.DB.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query, err := r.Queries.Get("create_relation.cypher")
		if err != nil {
			return nil, err
		}

		query = fmt.Sprintf(query, string(relation.Rel))

		params := map[string]any{
			"world_id":    worldID,
			"from_id":     relation.FromID,
			"to_id":       relation.To.ID,
			"description": relation.Description,
			"created_by":  relation.CreatedBy,
			"created_at":  relation.CreatedAt.Format(time.RFC3339),
			"updated_at":  relation.UpdatedAt.Format(time.RFC3339),
		}

		result, err := tx.Run(ctx, query, params)

		if err != nil {
			log.Println("err: ", err)
			return nil, fmt.Errorf("failed to create relation (%s) b/w nodes %s & %s\n", relation.Rel, relation.FromID, relation.To.ID)
		}

		return result.Single(ctx)
	})

	if err != nil {
		return nil, err
	}

	return relation, nil
}

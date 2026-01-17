package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (r *LoreRepository) CreateRelation(ctx context.Context, worldID string, relation model.Relation) (*model.LoreNode, *model.Relation, *model.LoreNode, error) {
	result, err := r.DB.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query, err := r.Queries.Get("create_relation.cypher")
		if err != nil {
			return nil, err
		}

		query = fmt.Sprintf(query, string(relation.Rel))

		params := map[string]any{
			"world_id":    worldID,
			"from_id":     relation.FromID,
			"to_id":       relation.ToID,
			"description": relation.Description,
			"created_by":  relation.CreatedBy,
			"created_at":  relation.CreatedAt.Format(time.RFC3339),
			"updated_at":  relation.UpdatedAt.Format(time.RFC3339),
		}

		result, err := tx.Run(ctx, query, params)

		if err != nil {
			log.Println("err: ", err)
			return nil, fmt.Errorf("failed to create relation (%s) b/w nodes %s & %s\n", relation.Rel, relation.FromID, relation.ToID)
		}

		return result.Single(ctx)
	})

	if err != nil {
		return nil, nil, nil, err
	}

	record := result.(*neo4j.Record)
	fromNode, ok := record.Get("from")
	if !ok {
		return nil, nil, nil, errors.New("missing 'from' node")
	}

	relValue, ok := record.Get("rel")
	if !ok {
		return nil, nil, nil, errors.New("missing 'rel' relationship")
	}

	toNode, ok := record.Get("to")
	if !ok {
		return nil, nil, nil, errors.New("missing 'to' node")
	}

	from := mapNode(fromNode.(neo4j.Node))
	to := mapNode(toNode.(neo4j.Node))
	rel := mapRelation(from, to, relValue.(neo4j.Relationship))

	return from, rel, to, nil
}

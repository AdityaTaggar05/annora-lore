package repository

import (
	"context"
	"errors"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (r *LoreRepository) CreateNode(ctx context.Context, node model.LoreNode) (*model.LoreNode, error) {
	result, err := r.DB.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query, err := r.Queries.Get("create_node.cypher")
		if err != nil {
			return nil, err
		}

		params := map[string]any{
			"id":           node.ID,
			"type":         string(node.Type),
			"name":         node.Name,
			"world_id":     node.WorldID,
			"created_by":   node.CreatedBy,
			"canon_status": string(node.CanonStatus),
			"created_at":   node.CreatedAt.Format(time.RFC3339),
			"updated_at":   node.UpdatedAt.Format(time.RFC3339),
		}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, errors.New("failed to create node: " + err.Error())
		}

		if !result.Next(ctx) {
			return nil, errors.New("node creation did not return result")
		}

		return *result.Record(), nil
	})

	if err != nil {
		return nil, err
	}

	return MapNode(result.(neo4j.Record)), err
}

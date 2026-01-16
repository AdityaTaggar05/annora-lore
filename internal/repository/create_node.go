package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (r *LoreRepository) CreateNode(ctx context.Context, node model.LoreNode) (*model.LoreNode, error) {
	result, err := r.DB.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
			MATCH (w:WorldNode {id: $world_id})
			CREATE (n:LoreNode {
				id: $id,
				type: $type,
				name: $name,
				world_id: $world_id,
				created_by: $created_by,
				canon_status: $canon_status,
				created_at: datetime($created_at),
				updated_at: datetime($updated_at),
				custom: $custom
			})
			CREATE (w)-[:CONTAINS]->(n)
			RETURN n
		`

		customMap, err := json.Marshal(node.Custom)
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
			"custom":       string(customMap),
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

	return MapNode(result.(neo4j.Record)), err
}

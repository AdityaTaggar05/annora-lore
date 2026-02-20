package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (r *LoreRepository) FetchNodeByID(ctx context.Context, worldID string, nodeID string) (*model.LoreNode, error) {
	result, err := r.DB.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query, err := r.Queries.Get("fetch_node_by_id.cypher")
		if err != nil {
			return nil, err
		}

		params := map[string]any{
			"world_id": worldID,
			"node_id": nodeID,
		}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			log.Println("err:", err)
			return nil, fmt.Errorf("failed to fetch node with id: %s", nodeID)
		}

		return result.Single(ctx)
	})

	if err != nil {
		return nil, err
	}

	record := result.(*neo4j.Record)

	node, ok := record.Get("node")
	if !ok {
		return nil, errors.New("missing node 'node'")
	}

	return mapNode(node.(neo4j.Node)), nil
} 

func (r *LoreRepository) FetchNodesByType(ctx context.Context, worldID string, nodeType []string) ([]*model.LoreNode, error) {
	result, err := r.DB.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query, err := r.Queries.Get("fetch_nodes_by_type.cypher")
		if err != nil {
			return nil, err
		}

		params := map[string]any{
			"world_id": worldID,
			"node_types": nodeType,
		}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		return result.Collect(ctx)
	})

	if err != nil {
		return nil, err
	}

	records := result.([]*neo4j.Record)
	nodes := make([]*model.LoreNode, len(records))

	for i, record := range records {
		node, ok := record.Get("n")
		if !ok {
			return nil, errors.New("missing node 'n'")
		}

		nodes[i] = mapNode(node.(neo4j.Node))
	}

	return nodes, nil
}

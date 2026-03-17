package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (r *LoreRepository) FetchAllRelationsOfNode(ctx context.Context, nodeID string) ([]*model.Relation, error) {
	result, err := r.DB.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query, err := r.Queries.Get("fetch_all_relations_of_node.cypher")
		if err != nil {
			return nil, err
		}

		params := map[string]any{
			"node_id": nodeID,
		}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch relations of node with id: %s", nodeID)
		}

		return result.Collect(ctx)
	})

	if err != nil {
		return nil, err
	}

	records := result.([]*neo4j.Record)

	relations := make([]*model.Relation, len(records))
	for i, record := range records {
		to, ok := record.Get("to")
		if !ok {
			return nil, errors.New("missing node 'to'")
		}

		rel, ok := record.Get("r")
		if !ok {
			return nil, errors.New("missing relationship 'r'")
		}

		relations[i] = mapRelation(rel.(neo4j.Relationship))
		relations[i].FromID = nodeID
		relations[i].To = mapNode(to.(neo4j.Node))
	}

	return relations, nil
}

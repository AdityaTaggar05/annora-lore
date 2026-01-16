package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func MapNode(record neo4j.Record) *model.LoreNode {
	node := record.Values[0].(neo4j.Node)
	props := node.Props

	var custom map[string]any
	err := json.Unmarshal([]byte(props["custom"].(string)), &custom)

	if err != nil {
		fmt.Println("err: ", err.Error())
		return nil
	}

	return &model.LoreNode{
		ID:          node.ElementId,
		Type:        model.NodeType(props["type"].(string)),
		Name:        props["name"].(string),
		CreatedBy:   props["created_by"].(string),
		CreatedAt:   props["created_at"].(time.Time),
		UpdatedAt:   props["updated_at"].(time.Time),
		CanonStatus: model.CanonStatus(props["canon_status"].(string)),
		WorldID:     props["world_id"].(string),
		Custom:      custom,
	}
}

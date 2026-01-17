package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AdityaTaggar05/annora-lore/internal/domain/model"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func mapNode(node neo4j.Node) *model.LoreNode {
	var custom map[string]any
	err := json.Unmarshal([]byte(node.Props["custom"].(string)), &custom)

	if err != nil {
		fmt.Println("err: ", err.Error())
		return nil
	}

	return &model.LoreNode{
		ID:          node.Props["id"].(string),
		Type:        model.NodeType(node.Props["type"].(string)),
		Name:        node.Props["name"].(string),
		CreatedBy:   node.Props["created_by"].(string),
		CreatedAt:   parseTime(node.Props["created_at"]),
		UpdatedAt:   parseTime(node.Props["updated_at"]),
		CanonStatus: model.CanonStatus(node.Props["canon_status"].(string)),
		WorldID:     node.Props["world_id"].(string),
		Custom:      custom,
	}
}

func mapRelation(from, to *model.LoreNode, rel neo4j.Relationship) *model.Relation {
	return &model.Relation{
		FromID:      from.ID,
		ToID:        to.ID,
		Rel:         model.RelationType(rel.Type),
		Description: rel.Props["description"].(string),
		CreatedBy:   rel.Props["created_by"].(string),
		CreatedAt:   parseTime(rel.Props["created_at"]),
		UpdatedAt:   parseTime(rel.Props["updated_at"]),
	}
}

func parseTime(v any) time.Time {
	switch t := v.(type) {
	case time.Time:
		return t
	case string:
		parsed, _ := time.Parse(time.RFC3339, t)
		return parsed
	default:
		return time.Time{}
	}
}

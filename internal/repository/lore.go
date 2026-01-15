package repository

import (
	"github.com/AdityaTaggar05/annora-lore/internal/infrastructure/neo4j"
)

type LoreRepository struct {
	DB *neo4j.Neo4jDB
}

func NewLoreRepository(db *neo4j.Neo4jDB) *LoreRepository {
	return &LoreRepository{DB: db}
}

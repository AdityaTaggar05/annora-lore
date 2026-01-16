package repository

import (
	"github.com/AdityaTaggar05/annora-lore/internal/infrastructure/neo4j"
)

type LoreRepository struct {
	DB      *neo4j.Neo4jDB
	Queries *QueryLoader
}

func NewLoreRepository(db *neo4j.Neo4jDB) (*LoreRepository, error) {
	ql, err := newQueryLoader("internal/repository/queries")
	if err != nil {
		return nil, err
	}

	return &LoreRepository{DB: db, Queries: ql}, nil
}

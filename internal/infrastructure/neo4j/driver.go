package neo4j

import (
	"context"
	"fmt"

	"github.com/AdityaTaggar05/annora-lore/internal/config"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

type Neo4jDB struct {
	Driver neo4j.Driver
	Config *config.Neo4jConfig
}

func NewNeo4jDB(ctx context.Context, cfg *config.Neo4jConfig) (*Neo4jDB, error) {
	driver, err := neo4j.NewDriver(
		cfg.URI,
		neo4j.BasicAuth(cfg.Username, cfg.Password, ""),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create neo4j driver: %w", err)
	}

	if err := driver.VerifyConnectivity(ctx); err != nil {
		return nil, fmt.Errorf("failed to verify neo4j connectivity: %w", err)
	}

	return &Neo4jDB{
		Driver: driver,
		Config: cfg,
	}, nil
}

func (db *Neo4jDB) Close(ctx context.Context) error {
	return db.Driver.Close(ctx)
}

func (db *Neo4jDB) ExecuteWrite(ctx context.Context, work neo4j.ManagedTransactionWork) (any, error) {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: db.Config.Database,
	})
	defer session.Close(ctx)

	return session.ExecuteWrite(ctx, work)
}

func (db *Neo4jDB) ExecuteRead(ctx context.Context, work neo4j.ManagedTransactionWork) (any, error) {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: db.Config.Database,
	})
	defer session.Close(ctx)

	return session.ExecuteRead(ctx, work)
}

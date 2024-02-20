package db

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/pkg/errors"
	"github.com/tj/go/env"
)

func ConnectToNeo4j() (neo4j.DriverWithContext, error) {
	uri := env.Get("NEO4J_CONN_STR")
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.NoAuth())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Neo4j driver: %v")
	}
	return driver, nil
}

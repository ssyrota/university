package graph_repo

import (
	"context"
	"distributed_systems_lab2/src/core"

	"github.com/elliotchance/pie/v2"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/pkg/errors"
)

func NewCityFactory(driver neo4j.DriverWithContext) *CityFactory {
	return &CityFactory{driver: driver}
}

type CityFactory struct {
	driver neo4j.DriverWithContext
}

type city struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

func (c *city) toCore() core.City {
	return *core.NewCity(uuid.MustParse(c.Id), c.Name)
}

func (f *CityFactory) ExistedInCvs() (*[]core.City, error) {
	ctx := context.Background()
	session := f.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	result, err := session.Run(ctx, existedInCvsCitiesQuery, map[string]any{})
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	var cities []city
	for result.Next(ctx) {
		record := result.Record().AsMap()
		c := city{
			Id:   record["c.id"].(string),
			Name: record["c.name"].(string),
		}
		cities = append(cities, c)
	}

	res := pie.Map(cities, func(c city) core.City { return c.toCore() })
	return &res, nil
}

const existedInCvsCitiesQuery = `
MATCH (:User)--(:Cv)--(:Job)--(c:City)
RETURN DISTINCT c.name, c.id
`

var _ core.CityFactory = new(CityFactory)

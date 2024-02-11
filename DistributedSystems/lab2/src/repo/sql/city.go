package sql

import (
	"context"
	"distributed_systems_lab2/src/core"
	"distributed_systems_lab2/src/db"

	"github.com/elliotchance/pie/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewCityFactory(db *db.SqlDB) *CityFactory {
	return &CityFactory{db: db}
}

type CityFactory struct {
	db *db.SqlDB
}

type city struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (c *city) toCore() core.City {
	return *core.NewCity(c.Id, c.Name)
}

var ctx = context.Background()

func (f *CityFactory) ExistedInCvs() (*[]core.City, error) {
	var cities []city
	if err := f.db.SelectContext(ctx, &cities, existedInCvsCitiesQuery); err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	res := pie.Map(cities, func(c city) core.City { return c.toCore() })
	return &res, nil
}

const existedInCvsCitiesQuery = `
SELECT
	id,
	name
FROM
	cities
WHERE EXISTS (
	SELECT 1 FROM cvs WHERE EXISTS (
		SELECT 1 from jobs WHERE jobs.cv_id = cvs.id AND job.city_id = cities.id
	)
)
`

var _ core.CityFactory = new(CityFactory)

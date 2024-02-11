package sql

import (
	"distributed_systems_lab2/src/core"
	"distributed_systems_lab2/src/db"

	"github.com/elliotchance/pie/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewHobbiesFactory(db *db.SqlDB) *HobbiesFactory {
	return &HobbiesFactory{db: db}
}

type HobbiesFactory struct {
	db *db.SqlDB
}

var _ core.HobbiesFactory = new(HobbiesFactory)

type hobby struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (c *hobby) toCore() core.Hobby {
	return core.NewHobby(c.Name, c.Id)
}

func (f *HobbiesFactory) ByUsersInCity(city core.City) (*[]core.Hobby, error) {
	var hobbies []hobby
	query, args := f.byUsersInCityQuery(city.ID())
	if err := f.db.SelectContext(ctx, &hobbies, query, args...); err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	res := pie.Map(hobbies, func(c hobby) core.Hobby { return c.toCore() })
	return &res, nil
}

func (f *HobbiesFactory) byUsersInCityQuery(cityID uuid.UUID) (string, []any) {
	query, args, _ := sqlx.Named(hobbiesByCityQueryT, map[string]any{
		"city_id": cityID.String(),
	})
	return f.db.Rebind(query), args
}

func (f *HobbiesFactory) ExistedInCvs() (*[]core.Hobby, error) {
	var hobbies []hobby
	if err := f.db.SelectContext(ctx, &hobbies, existedInCvsHobbiesQuery); err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	res := pie.Map(hobbies, func(c hobby) core.Hobby { return c.toCore() })
	return &res, nil
}

const hobbiesByCityQueryT = `
SELECT
	id,
	name
FROM
	hobbies
WHERE EXISTS (
	SELECT 1 from cvs_hobbies WHERE cvs_hobbies.cv_id IN (
		SELECT id FROM cvs WHERE EXISTS (
			SELECT 1 from jobs WHERE jobs.cv_id = cvs.id AND job.city_id=:city_id
		)
	) AND cvs_hobbies.hobby_id = hobbies.id
)
`

const existedInCvsHobbiesQuery = `
SELECT
	id,
	name
FROM
	hobbies
WHERE EXISTS (
	SELECT 1 FROM cvs WHERE EXISTS (
		SELECT 1 from cvs_hobbies WHERE cvs_hobbies.cv_id = cvs.id AND cvs_hobbies.hobby_id = hobbies.id
	)
)
`

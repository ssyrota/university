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

func (f *HobbiesFactory) ByUsersInCity(city string) (*[]core.Hobby, error) {
	var hobbies []hobby
	query, args := f.byUsersInCityQuery(city)
	if err := f.db.SelectContext(ctx, &hobbies, query, args...); err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	res := pie.Map(hobbies, func(c hobby) core.Hobby { return c.toCore() })
	return &res, nil
}

func (f *HobbiesFactory) byUsersInCityQuery(city string) (string, []any) {
	query, args, _ := sqlx.Named(hobbiesByCityQueryT, map[string]any{
		"city": city,
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
WHERE id IN (
	SELECT hobby_id from cvs_hobbies WHERE cv_id IN (
		SELECT id FROM cvs cv WHERE
	cv.id IN (
		SELECT cv_id FROM jobs job WHERE job.city_id IN (SELECT id FROM cities WHERE name=:city)
		) 
	)
)
`

const existedInCvsHobbiesQuery = `
SELECT
	id,
	name
FROM
	hobbies
WHERE id IN (SELECT hobby_id from cvs_hobbies WHERE cv_id IN (SELECT id from cvs))
`

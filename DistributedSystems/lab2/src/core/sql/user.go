package sql

import (
	"distributed_systems_lab2/src/core"
	"distributed_systems_lab2/src/db"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewUsersFactory(db *db.SqlDB) *UsersFactory {
	return &UsersFactory{db: db}
}

type UsersFactory struct {
	db *db.SqlDB
}

type user struct {
	login    string `db:"login"`
	password string `db:"password"`
}

func (c *user) toCore(f *UsersFactory) *core.User {
	return core.NewUser(func() (*core.Cv, error) {
		return nil, nil
	}, c.login, c.password)
}

func (f *UsersFactory) Find(login string) (*core.User, error) {
	var user user
	query, args := f.findUserQuery(login)
	if err := f.db.SelectContext(ctx, &user, query, args); err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	return user.toCore(f), nil
}

func (f *UsersFactory) findUserQuery(login string) (string, []any) {
	query, args, _ := sqlx.Named(userQueryT, map[string]any{
		"login": login,
	})
	return f.db.Rebind(query), args
}

const userQueryT = `
SELECT login, password FROM users WHERE login = :login
`

func (*UsersFactory) GroupByWorkedCompany() (map[string][]core.User, error) {
	panic("unimplemented")
}

var _ core.UsersFactory = new(UsersFactory)

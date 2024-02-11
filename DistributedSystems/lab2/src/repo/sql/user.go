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
	Login    string `db:"login"`
	Password string `db:"password"`
}

func (c *user) toCore(f *UsersFactory) *core.User {
	return core.NewUser(func() (*core.Cv, error) {
		return nil, nil
	}, c.Login, c.Password)
}

func (f *UsersFactory) Find(login string) (*core.User, error) {
	var users []user
	query, args := f.findUserQuery(login)
	if err := f.db.SelectContext(ctx, &users, query, args...); err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	return users[0].toCore(f), nil
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

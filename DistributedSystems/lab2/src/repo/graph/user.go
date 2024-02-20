package graph_repo

import (
	"context"
	"distributed_systems_lab2/src/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/pkg/errors"
)

func NewUsersFactory(driver neo4j.DriverWithContext) *UsersFactory {
	return &UsersFactory{driver: driver}
}

type UsersFactory struct {
	driver neo4j.DriverWithContext
}

func (c *user) toCore(f *UsersFactory) *core.User {
	return core.NewUser(func() (*core.Cv, error) {
		cv := c.Cv.toDomain()
		return &cv, nil
	}, c.Login, c.Password)
}

func (f *UsersFactory) Find(login string) (*core.User, error) {
	ctx := context.Background()
	session := f.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	result, err := session.Run(ctx, findUserQuery, map[string]any{})
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	dbUsers, err := result.Collect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	if len(dbUsers) == 0 {
		return nil, errors.New("user not found")
	}
	dbUser := dbUsers[0].AsMap()
	dbLogin, dbPassword := dbUser["u.login"].(string), dbUser["u.password"].(string)
	return core.NewUser(func() (*core.Cv, error) {
		return nil, nil
	}, dbLogin, dbPassword), nil
}
func (f *UsersFactory) GroupByWorkedCompany() (map[string][]core.User, error) {
	ctx := context.Background()
	session := f.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	result, err := session.Run(ctx, findUserByCompanyQuery, map[string]any{})
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	dbRes, err := result.Collect(ctx)
	domainRes := make(map[string][]core.User)
	for _, rawCompany := range dbRes {
		company := rawCompany.AsMap()
		users := []core.User{}
		dbUsers := company["users"].([]any)
		for _, dbUser := range dbUsers {
			user := dbUser.(map[string]any)
			users = append(users, *core.NewUser(func() (*core.Cv, error) {
				return nil, nil
			}, user["login"].(string), user["password"].(string)))
		}
		domainRes[company["c.name"].(string)] = users
	}

	return domainRes, nil
}

var _ core.UsersFactory = new(UsersFactory)

const findUserQuery = `
MATCH (u:User {login: $login})
RETURN u.login, u.password
`

const findUserByCompanyQuery = `
MATCH (u:User)--(:Cv)--(:Job)--(c:Company)
RETURN c.name, collect({login:u.login, password:u.password}) AS users
`

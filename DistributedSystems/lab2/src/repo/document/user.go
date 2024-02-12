package documented_repo

import (
	"context"
	"distributed_systems_lab2/src/core"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUsersFactory(db *mongo.Database) *UsersFactory {
	return &UsersFactory{db: db.Collection("users_with_cvs")}
}

type UsersFactory struct {
	db *mongo.Collection
}

func (c *user) toCore(f *UsersFactory) *core.User {
	return core.NewUser(func() (*core.Cv, error) {
		cv := c.Cv.toDomain()
		return &cv, nil
	}, c.Login, c.Password)
}

func (f *UsersFactory) Find(login string) (*core.User, error) {
	var dbUser user
	if err := f.db.FindOne(context.Background(), bson.M{"login": login}).Decode(&dbUser); err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	return dbUser.toCore(f), nil
}
func (*UsersFactory) GroupByWorkedCompany() (map[string][]core.User, error) {
	panic("unimplemented")
}

var _ core.UsersFactory = new(UsersFactory)

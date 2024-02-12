package documented_repo

import (
	"context"
	"distributed_systems_lab2/src/core"

	"github.com/elliotchance/pie/v2"
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
func (f *UsersFactory) GroupByWorkedCompany() (map[string][]core.User, error) {
	pipeline := bson.A{
		bson.D{{"$unwind", bson.D{{"path", "$cv.jobs"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$cv.jobs.company"},
					{"data", bson.D{{"$addToSet", "$$ROOT"}}},
				},
			},
		},
	}
	cursor, err := f.db.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	type res struct {
		company string `bson:"_id"`
		data    []user `bson:"data"`
	}
	var dbRes []res
	if err := cursor.All(ctx, &dbRes); err != nil {
		return nil, err
	}
	domainRes := make(map[string][]core.User)
	for _, dbResItem := range dbRes {
		domainRes[dbResItem.company] = pie.Map(dbResItem.data, func(u user) core.User {
			return *u.toCore(f)
		})
	}
	return domainRes, nil
}

var _ core.UsersFactory = new(UsersFactory)

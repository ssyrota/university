package documented_repo

import (
	"context"
	"distributed_systems_lab2/src/core"

	"github.com/elliotchance/pie/v2"
	"github.com/pkg/errors"
	"github.com/tj/go/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHobbiesFactory(db *mongo.Database) *HobbiesFactory {
	return &HobbiesFactory{db: db.Collection(env.Get("MONGO_COLLECTION"))}
}

type HobbiesFactory struct {
	db *mongo.Collection
}

var _ core.HobbiesFactory = new(HobbiesFactory)

// Springdale
func (f *HobbiesFactory) ByUsersInCity(city string) (*[]core.Hobby, error) {
	// If you are reading this and want to reuse in a more performant way use nested indexes
	pipeline := bson.A{
		bson.D{{"$unwind", bson.D{{"path", "$cv.jobs"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$cv.jobs.city.name"},
					{"hobbies", bson.D{{"$addToSet", "$cv.hobbies"}}},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$hobbies"}}}},
		bson.D{{"$unwind", bson.D{{"path", "$hobbies"}}}},
		bson.D{{"$match", bson.D{{"_id", city}}}},
	}
	cursor, err := f.db.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	type res struct {
		Hobbies hobby `bson:"hobbies"`
	}
	var dbRes []res
	if err := cursor.All(ctx, &dbRes); err != nil {
		return nil, err
	}
	domainRes := pie.Map(dbRes, func(r res) core.Hobby {
		return core.Hobby(r.Hobbies.toDomain())
	})
	return &domainRes, nil
}

func (f *HobbiesFactory) ExistedInCvs() (*[]core.Hobby, error) {
	// If you are reading this and want to reuse in a more performant way use nested indexes
	pipeline := bson.A{
		bson.D{{Key: "$unwind", Value: bson.D{{"path", "$cv.jobs"}}}},
		bson.D{
			{Key: "$group",
				Value: bson.D{
					{Key: "_id", Value: "$cv.jobs.city.name"},
					{Key: "hobbies", Value: bson.D{{"$addToSet", "$cv.hobbies"}}},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: bson.D{{"path", "$hobbies"}}}},
		bson.D{{Key: "$unwind", Value: bson.D{{"path", "$hobbies"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", primitive.Null{}},
					{"deduped", bson.D{{"$addToSet", "$hobbies"}}},
				},
			},
		},
	}
	cursor, err := f.db.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "aggregation failed to start")
	}
	type res struct {
		Deduped []hobby `bson:"deduped"`
	}
	var dbRes []res
	if err := cursor.All(context.Background(), &dbRes); err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	domainRes := pie.Map(dbRes[0].Deduped, func(r hobby) core.Hobby {
		return r.toDomain()
	})
	return &domainRes, nil
}

package documented_repo

import (
	"distributed_systems_lab2/src/core"

	"github.com/elliotchance/pie/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHobbiesFactory(db *mongo.Database) *HobbiesFactory {
	return &HobbiesFactory{db: db.Collection("users_with_cvs")}
}

type HobbiesFactory struct {
	db *mongo.Collection
}

var _ core.HobbiesFactory = new(HobbiesFactory)

// Springdale
func (f *HobbiesFactory) ByUsersInCity(city string) (*[]core.Hobby, error) {
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
		hobbies hobby `bson:"hobbies"`
	}
	var dbRes []res
	if err := cursor.All(ctx, &dbRes); err != nil {
		return nil, err
	}
	domainRes := pie.Map(dbRes, func(r res) core.Hobby {
		return core.Hobby(r.hobbies.toDomain())
	})
	return &domainRes, nil
}

func (f *HobbiesFactory) ExistedInCvs() (*[]core.Hobby, error) {
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
		bson.D{{Key: "$unwiValue: nd", Value: bson.D{{"path", "$hobbies"}}}},
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
		return nil, err
	}
	type res struct {
		deduped []hobby `bson:"deduped"`
	}
	var dbRes res
	if err := cursor.Decode(dbRes); err != nil {
		return nil, err
	}
	domainRes := pie.Map(dbRes.deduped, func(r hobby) core.Hobby {
		return r.toDomain()
	})
	return &domainRes, nil
}

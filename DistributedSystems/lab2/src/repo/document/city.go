package documented_repo

import (
	"context"
	"distributed_systems_lab2/src/core"

	"github.com/elliotchance/pie/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCityFactory(db *mongo.Database) *CityFactory {
	return &CityFactory{db: db.Collection("users_with_cvs")}
}

type CityFactory struct {
	db *mongo.Collection
}

var ctx = context.Background()

func (f *CityFactory) ExistedInCvs() (*[]core.City, error) {
	pipeline := bson.A{
		bson.D{{"$unwind", bson.D{{"path", "$cv.jobs"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$cv.jobs.city"},
					{"data", bson.D{{"$addToSet", primitive.Null{}}}},
				},
			},
		},
	}
	cursor, err := f.db.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	type res struct {
		data city `bson:"data"`
	}
	var dbRes []res
	if err := cursor.All(ctx, &dbRes); err != nil {
		return nil, err
	}
	domainRes := pie.Map(dbRes, func(r res) core.City {
		return *r.data.toDomain()
	})
	return &domainRes, nil
}

var _ core.CityFactory = new(CityFactory)

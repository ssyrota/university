package documented_repo

import (
	"context"
	"distributed_systems_lab2/src/core"

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
	var cities []city
}

var _ core.CityFactory = new(CityFactory)

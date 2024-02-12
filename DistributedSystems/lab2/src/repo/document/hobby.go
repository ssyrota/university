package documented_repo

import (
	"distributed_systems_lab2/src/core"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewHobbiesFactory(db *mongo.Database) *HobbiesFactory {
	return &HobbiesFactory{db: db.Collection("users_with_cvs")}
}

type HobbiesFactory struct {
	db *mongo.Collection
}

var _ core.HobbiesFactory = new(HobbiesFactory)

func (f *HobbiesFactory) ByUsersInCity(city string) (*[]core.Hobby, error) {
	var hobbies []hobby
}
func (f *HobbiesFactory) ExistedInCvs() (*[]core.Hobby, error) {
	var hobbies []hobby
}

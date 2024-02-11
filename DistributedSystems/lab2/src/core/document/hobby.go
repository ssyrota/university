package sql

import "distributed_systems_lab2/src/core"

type HobbiesFactory struct {
}

var _ core.HobbiesFactory = new(HobbiesFactory)

// ByUsersInCity implements core.HobbiesFactory.
func (*HobbiesFactory) ByUsersInCity(city core.City) (*[]core.Hobby, error) {
	panic("unimplemented")
}

// ExistedInCvs implements core.HobbiesFactory.
func (*HobbiesFactory) ExistedInCvs() (*[]core.Hobby, error) {
	panic("unimplemented")
}

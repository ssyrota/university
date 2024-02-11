package sql

import "distributed_systems_lab2/src/core"

type CityFactory struct {
}

// ExistedInCvs implements core.CityFactory.
func (*CityFactory) ExistedInCvs() (*[]core.City, error) {
	panic("unimplemented")
}

var _ core.CityFactory = new(CityFactory)

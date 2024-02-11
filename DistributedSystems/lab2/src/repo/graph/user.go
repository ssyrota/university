package sql

import "distributed_systems_lab2/src/core"

type UsersFactory struct {
}

// Find implements core.UsersFactory.
func (*UsersFactory) Find(login string) (*core.User, error) {
	panic("unimplemented")
}

// GroupByWorkedCompany implements core.UsersFactory.
func (*UsersFactory) GroupByWorkedCompany() (map[string][]core.User, error) {
	panic("unimplemented")
}

var _ core.UsersFactory = new(UsersFactory)

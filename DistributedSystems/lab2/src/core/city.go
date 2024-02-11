package core

import "github.com/google/uuid"

type City struct {
	Id   uuid.UUID
	Name string
}

func (c *City) ID() uuid.UUID {
	return c.Id
}

func NewCity(id uuid.UUID, name string) *City {
	return &City{Id: id, Name: name}
}

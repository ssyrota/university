package core

import "github.com/google/uuid"

func NewHobby(name string, id uuid.UUID) Hobby {
	return Hobby{id, name}
}

type Hobby struct {
	Id   uuid.UUID
	Name string
}

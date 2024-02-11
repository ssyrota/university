package core

import (
	"github.com/google/uuid"
)

func NewCv(id uuid.UUID, hobbies LazyData[[]Hobby], jobHistory LazyData[[]Job]) Cv {
	return Cv{id, hobbies, jobHistory}
}

type Cv struct {
	Id uuid.UUID
	// *-*
	Hobbies LazyData[[]Hobby]
	// 1-*
	JobHistory LazyData[[]Job]
}

package core

import (
	"time"

	"github.com/google/uuid"
)

func NewJob(id uuid.UUID, from time.Time, to time.Time, company string, city City) Job {
	return Job{id, from, to, company, city}
}

type Job struct {
	Id uuid.UUID

	From    time.Time
	To      time.Time
	Company string
	City    City
}

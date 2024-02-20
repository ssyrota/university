package graph_repo

import (
	"distributed_systems_lab2/src/core"
	"time"

	"github.com/google/uuid"
)

type user struct {
	Login    string
	Password string
	Cv       cv
}

type cv struct {
	Id         string
	Hobbies    []hobby
	JobHistory []job
}

func (cv *cv) toDomain() core.Cv {
	return core.NewCv(
		uuid.MustParse(cv.Id),
		func() (*[]core.Hobby, error) {
			hobbies := make([]core.Hobby, len(cv.Hobbies))
			for i, h := range cv.Hobbies {
				hobbies[i] = h.toDomain()
			}
			return &hobbies, nil
		},
		func() (*[]core.Job, error) {
			jobs := make([]core.Job, len(cv.JobHistory))
			for i, j := range cv.JobHistory {
				jobs[i] = j.toDomain()
			}
			return &jobs, nil
		},
	)
}

type hobby struct {
	Id   string
	Name string
}

func (hobby *hobby) toDomain() core.Hobby {
	return core.NewHobby(hobby.Name, uuid.MustParse(hobby.Id))
}

type job struct {
	Id      string
	From    time.Time
	To      time.Time
	Company string
	City    city
}

func (job *job) toDomain() core.Job {
	return core.NewJob(
		uuid.MustParse(job.Id),
		job.From,
		job.To,
		job.Company,
		job.City.toCore(),
	)

}

package documented_repo

import (
	"distributed_systems_lab2/src/core"
	"time"

	"github.com/google/uuid"
)

type user struct {
	Login    string `bson:"login"`
	Password string `bson:"password"`
	Cv       cv     `bson:"cv"`
}

type cv struct {
	Id string `bson:"login"`
	// *-*
	Hobbies []hobby
	// 1-*
	JobHistory []job
}

func (cv *cv) toDomain() core.Cv {
	return core.NewCv(
		uuid.MustParse(cv.Id),
		// preload lazy data
		func() (*[]core.Hobby, error) {
			hobbies := make([]core.Hobby, len(cv.Hobbies))
			for i, h := range cv.Hobbies {
				hobbies[i] = h.toDomain()
			}
			return &hobbies, nil
		},
		// preload lazy data
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
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

func (hobby *hobby) toDomain() core.Hobby {
	return core.NewHobby(hobby.Name, uuid.MustParse(hobby.Id))
}

type job struct {
	Id      string    `bson:"id"`
	From    time.Time `bson:"from"`
	To      time.Time `bson:"to"`
	Company string    `bson:"company"`
	City    city      `bson:"city"`
}

func (job *job) toDomain() core.Job {
	return core.NewJob(
		uuid.MustParse(job.Id),
		job.From,
		job.To,
		job.Company,
		*job.City.toDomain(),
	)

}

type city struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

func (city *city) toDomain() *core.City {
	return core.NewCity(uuid.MustParse(city.Id), city.Name)
}

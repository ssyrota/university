package core

type UsersFactory interface {
	Find(login string) (*User, error)
	GroupByWorkedCompany() (map[string][]User, error)
}

type HobbiesFactory interface {
	ExistedInCvs() (*[]Hobby, error)
	ByUsersInCity(city string) (*[]Hobby, error)
}

type CityFactory interface {
	ExistedInCvs() (*[]City, error)
}

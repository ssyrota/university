package core

import (
	"github.com/pkg/errors"
)

func (u *User) Cv() (*Cv, error) {
	cv, err := u.cv()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cv")
	}
	return cv, nil
}

type User struct {
	Login    string
	Password Password `json:"-"`
	cv       LazyData[Cv]
}

func NewUser(cv LazyData[Cv], login string, password Password) *User {
	return &User{login, password, cv}
}

type Password = string

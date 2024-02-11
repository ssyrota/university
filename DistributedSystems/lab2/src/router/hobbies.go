package router

import (
	"distributed_systems_lab2/src/core"
	"distributed_systems_lab2/src/ginres"

	"github.com/gin-gonic/gin"
)

type hobbies struct {
	hobbiesFactory core.HobbiesFactory
}

func (u *hobbies) AllInCvs(c *gin.Context) {
	var query loginUserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ginres.NewValidationErr(c, err).Reply()
		return
	}
	hobbies, err := u.hobbiesFactory.ExistedInCvs()
	if err != nil {
		ginres.NewInternalServerError(c, err).Reply()
		return
	}
	c.JSON(200, hobbies)
}

type allByCity struct {
	City string `json:"city"`
}

func (u *hobbies) AllByCity(c *gin.Context) {
	var query allByCity
	if err := c.ShouldBindQuery(&query); err != nil {
		ginres.NewValidationErr(c, err).Reply()
		return
	}
	hobbies, err := u.hobbiesFactory.ByUsersInCity(query.City)
	if err != nil {
		ginres.NewInternalServerError(c, err).Reply()
		return
	}
	c.JSON(200, hobbies)
}

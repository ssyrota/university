package router

import (
	"distributed_systems_lab2/src/core"
	"distributed_systems_lab2/src/ginres"

	"github.com/gin-gonic/gin"
)

type citties struct {
	cittiesFactory core.CityFactory
}

func (u *citties) AllInCvs(c *gin.Context) {
	var query loginUserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ginres.NewValidationErr(c, err).Reply()
		return
	}
	cities, err := u.cittiesFactory.ExistedInCvs()
	if err != nil {
		ginres.NewInternalServerError(c, err).Reply()
		return
	}
	c.JSON(200, cities)
}

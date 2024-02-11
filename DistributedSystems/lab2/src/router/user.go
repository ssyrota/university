package router

import (
	"distributed_systems_lab2/src/core"
	"distributed_systems_lab2/src/ginres"

	"github.com/gin-gonic/gin"
)

type user struct {
	userFactory core.UsersFactory
}

type loginUserQuery struct {
	Login string `json:"login"`
}

func (u *user) Cv(c *gin.Context) {
	var query loginUserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ginres.NewValidationErr(c, err).Reply()
		return
	}
	user, err := u.userFactory.Find(query.Login)
	if err != nil {
		ginres.NewInternalServerError(c, err).Reply()
		return
	}
	cv, err := user.Cv()
	if err != nil {
		ginres.NewInternalServerError(c, err).Reply()
		return
	}
	c.JSON(200, cv)
}

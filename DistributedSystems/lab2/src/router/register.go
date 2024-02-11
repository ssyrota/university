package router

import (
	"distributed_systems_lab2/src/core"

	"github.com/gin-gonic/gin"
)

func Register(group gin.IRouter, userFactory core.UsersFactory, cityFactory core.CityFactory, hobbiesFactory core.HobbiesFactory) {
	userController := user{userFactory: userFactory}
	group.GET("/cv", userController.Cv)
}

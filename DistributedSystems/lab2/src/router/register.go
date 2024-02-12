package router

import (
	"distributed_systems_lab2/src/core"

	"github.com/gin-gonic/gin"
)

func Register(group gin.IRouter, userFactory core.UsersFactory, cityFactory core.CityFactory, hobbiesFactory core.HobbiesFactory) {
	userController := user{userFactory: userFactory}
	group.GET("/user/cv", userController.Cv)
	group.GET("/user/hobbies", userController.Hobbies)
	group.GET("/user/company-grouped", userController.ByCompany)

	hobbiesController := hobbies{hobbiesFactory: hobbiesFactory}
	group.GET("/hobbies", hobbiesController.AllInCvs)
	group.GET("/hobbies/by-city", hobbiesController.AllByCity)

	citiesController := citties{cittiesFactory: cityFactory}
	group.GET("/cities", citiesController.AllInCvs)
}

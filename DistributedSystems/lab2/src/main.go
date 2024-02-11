package main

import (
	"distributed_systems_lab2/src/core/sql"
	"distributed_systems_lab2/src/db"
	"distributed_systems_lab2/src/router"
	"distributed_systems_lab2/src/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/tj/go/env"
)

func main() {
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	{
		sqlDb, err := db.NewWithPg(env.Get("GOOSE_DBSTRING"))
		if err != nil {
			panic(err)
		}
		sqlUsers := sql.NewUsersFactory(sqlDb)
		sqlCity := sql.NewCityFactory(sqlDb)
		sqlHobbies := sql.NewHobbiesFactory(sqlDb)
		router.Register(server.Group("/sql"), sqlUsers, sqlCity, sqlHobbies)
	}
	server.Run(":3001")
}

package main

import (
	"distributed_systems_lab2/src/db"
	documented_repo "distributed_systems_lab2/src/repo/document"
	graph_repo "distributed_systems_lab2/src/repo/graph"
	"distributed_systems_lab2/src/repo/sql"
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
		users := sql.NewUsersFactory(sqlDb)
		city := sql.NewCityFactory(sqlDb)
		hobbies := sql.NewHobbiesFactory(sqlDb)
		router.Register(server.Group("/sql"), users, city, hobbies)
	}

	{
		documentedDb, err := db.NewDocumented(env.Get("MONGO_CONNECTION_STRING"), env.Get("MONGO_DB"))
		if err != nil {
			panic(err)
		}
		users := documented_repo.NewUsersFactory(documentedDb)
		city := documented_repo.NewCityFactory(documentedDb)
		hobbies := documented_repo.NewHobbiesFactory(documentedDb)
		router.Register(server.Group("/documented"), users, city, hobbies)
	}

	{
		neo4j, err := db.ConnectToNeo4j()
		if err != nil {
			panic(err)
		}
		city := graph_repo.NewCityFactory(neo4j)
		user := graph_repo.NewUsersFactory(neo4j)
		router.Register(server.Group("/graph"), user, city, nil)
	}
	server.Run(":3001")
}

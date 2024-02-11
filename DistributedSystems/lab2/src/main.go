package main

import (
	"distributed_systems_lab2/src/http/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
}

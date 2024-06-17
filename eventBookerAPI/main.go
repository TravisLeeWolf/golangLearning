package main

import (
	"eventbooker.com/db"
	"eventbooker.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run("localhost:8080")
}

package main

import (
	"server/db"
	route "server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	route.RegisterRoutes(server)

	server.Run(":8080")

}

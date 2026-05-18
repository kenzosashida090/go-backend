package main

import (
	"db-go.com/api/db"
	"db-go.com/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	defer db.DB.Close()
	server.Run(":3001")
}

package main

import (
	"interactive-map/config"
	"interactive-map/database"
	"interactive-map/routes"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	routes.SetupRoutes(r)

	r.Run(":8080")
}

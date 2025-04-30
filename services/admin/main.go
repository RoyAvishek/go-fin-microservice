package main

import (
	"log"

	"go-fin-microservice/services/admin/internal/config"
	"go-fin-microservice/services/admin/internal/database"
	"go-fin-microservice/services/admin/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the database
	database.InitializeDatabase()
	database.MigrateDB()
	database.SeedDB()
	// Confirm successful initialization
	database.GetDB()
	log.Println("Database connection is ready:")

	config.InitConfig()

	port := config.Config.GetString("app.port")
	env := config.Config.GetString("app.environment")

	router := gin.Default()

	router.GET("/health", getHealthCheck)
	routes.RegisterRoutes(router)

	log.Println("Admin service is running on port: ", port)
	log.Printf("Started Admin service in %s mode", env)

	if err := router.Run(port); err != nil {
		log.Panic("Failed to start the server due to:", err)
	}
}

func getHealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Admin service is healthy",
	})
}

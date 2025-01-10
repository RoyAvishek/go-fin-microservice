package main

import (
	"log"

	"go-fin-microservice/services/admin/internal/config"
	"go-fin-microservice/services/admin/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the database
	database.InitializeDatabase()

	// Confirm successful initialization
	db := database.GetDB()
	log.Println("Database connection is ready:", db)

	config.InitConfig()

	port := config.Config.GetString("app.port")
	env := config.Config.GetString("app.environment")

	router := gin.Default()

	router.GET("/health", getHealthCheck)

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

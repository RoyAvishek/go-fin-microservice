package database

import (
	"fmt"
	"log"
	"os"

	"go-fin-microservice/services/admin/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global database instance

// InitializeDatabase sets up the database connection and stores it in the global DB variable.
func InitializeDatabase() {
	// Connection string for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	var err error

	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	// Log the success message
	log.Println("Database connection successfully established!")
}

// GetDB returns the global DB instance for use in other packages.
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized. Call InitializeDatabase first.")
	}
	return DB
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
}

func SeedDB() {
	// Create roles
	adminRole := models.Role{Name: "admin"}
	userRole := models.Role{Name: "user"}

	// Save roles to the database
	DB.Create(&adminRole)
	DB.Create(&userRole)

	// Create users
	user1 := models.User{
		Username: "john",
		Email:    "john@example.com",
		Password: "password",
	}
	user2 := models.User{
		Username: "jane",
		Email:    "jane@example.com",
		Password: "password",
	}

	// Save users to the database
	DB.Create(&user1)
	DB.Create(&user2)

	// Associate roles with users
	DB.Model(&user1).Association("Roles").Append(&adminRole)
	DB.Model(&user2).Association("Roles").Append(&userRole)

	// Log the success message
	log.Println("Database seeding completed successfully!")
}

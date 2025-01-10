package database

import (
	"log"
	"time"

	"go-fin-microservice/services/admin/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global database instance

// InitializeDatabase sets up the database connection and stores it in the global DB variable.
func InitializeDatabase() {
	// Connection string for PostgreSQL
	dsn := "host=localhost user=postgres password=go-pg-password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	var err error

	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Log the success message
	log.Println("Database connection successfully established!")
}

// GetDB returns the global DB instance for use in other packages.
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatalf("Database connection is not initialized. Call InitializeDatabase first.")
	}
	return DB
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func SeedDB() {
	// Create roles
	adminRole := models.Role{Name: "admin"}
	userRole := models.Role{Name: "user"}

	// Create permissions
	readPermission := models.Permission{Name: "read", Description: "Read permission"}
	writePermission := models.Permission{Name: "write", Description: "Write permission"}

	// Assign permissions to roles
	adminRole.Permissions = []models.Permission{readPermission, writePermission}
	userRole.Permissions = []models.Permission{readPermission}

	// Create users
	john := models.User{
		Username:  "john",
		Email:     "john@example.com",
		Password:  "xxxxxxxxx",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Roles:     []models.Role{adminRole},
	}

	jane := models.User{
		Username:  "jane",
		Email:     "jane@example.com",
		Password:  "xxxxxxxxx",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Roles:     []models.Role{userRole},
	}

	// Save to the database
	DB.Create(&john)
	DB.Create(&jane)
}

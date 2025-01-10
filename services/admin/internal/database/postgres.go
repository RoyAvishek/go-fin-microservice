package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global database instance

// InitializeDatabase sets up the database connection and stores it in the global DB variable.
func InitializeDatabase() {
	// Connection string for PostgreSQL
	// Replace with your actual credentials and database name
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

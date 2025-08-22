package main

import (
	"log"
	"service-b/config"
	"service-b/database/migrations"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	// Loan DB config from env
	databaseConfig := config.LoadDatabaseConfig()

	// Connect to DB
	db, err := config.ConnectDB(databaseConfig)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Run database migrations
	migrations.RunDatabaseMigrations(db, databaseConfig.DBName)

	// Create Echo instance
	e := echo.New()

	// Start server
	port := "8080"
	log.Printf("🚀 Starting Echo server on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal("❌ Echo server failed:", err)
	}
}

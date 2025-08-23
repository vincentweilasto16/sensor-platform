package main

import (
	"log"
	"service-b/database/migrations"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"service-b/internal/config"
	"service-b/internal/controller"
	repo "service-b/internal/repository/mysql"
	"service-b/internal/router"
	"service-b/internal/service"
	"service-b/internal/validator"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Load DB config from env
	databaseConfig := config.LoadDatabaseConfig()

	// Connect to DB
	db, err := config.ConnectDB(databaseConfig)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Run migrations
	migrations.RunDatabaseMigrations(db, databaseConfig.DBName)

	// Initialize repository
	dbRepo := repo.New(db)

	// Initialize services
	sensorService := service.NewSensorService(dbRepo)

	// Initialize controllers
	sensorController := controller.NewSensorController(sensorService)

	ctrl := &controller.Controllers{
		SensorController: sensorController,
	}

	// Create Echo instance
	e := echo.New()

	// Set Validator
	e.Validator = validator.New()

	// Setup Echo router
	router.NewRouter(e, ctrl)

	// Start server
	port := "8080"
	log.Printf("🚀 Starting Echo server on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal("❌ Echo server failed:", err)
	}
}

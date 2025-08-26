package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"service-b/database/migrations"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"service-b/internal/config"
	"service-b/internal/controller"
	"service-b/internal/handler/kafka"
	messaging "service-b/internal/messaging/kafka"
	repo "service-b/internal/repository/mysql"
	"service-b/internal/router"
	"service-b/internal/service"
	"service-b/internal/validator"
)

func main() {
	ctx := context.Background()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Load App Config from env
	appConfig := config.LoadAppConfig()

	// Load DB config from env
	databaseConfig := config.LoadDatabaseConfig()

	// Load Kafka config from env
	kafkaConfig, err := config.LoadKafkaConfig()
	if err != nil {
		log.Fatalf("❌ Failed to get kafka config: %v", err)
	}

	// Initialize kafka consumer
	kafkaConsumer := messaging.NewKafkaConsumer(kafkaConfig)

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

	// Initialize Kafka consumer handler
	sensorHandler := kafka.NewSensorHandler(sensorService)

	// Listen to Kafka messages in a goroutine
	go func() {
		if err := kafkaConsumer.Consume(ctx, func(key, value []byte) error {
			return sensorHandler.ProcessSensorEvent(ctx, value)
		}); err != nil {
			log.Printf("❌ Kafka consumer error: %v", err)
		}
	}()

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Println("⚠️ Shutting down gracefully...")
		kafkaConsumer.Close()
		os.Exit(0)
	}()

	// Start server
	log.Printf("🚀 Starting Echo server on port %s", appConfig.AppPort)
	if err := e.Start(":" + appConfig.AppPort); err != nil {
		log.Fatal("❌ Echo server failed:", err)
	}
}

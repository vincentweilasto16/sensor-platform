package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"service-a/internal/config"
	"service-a/internal/controller"
	"service-a/internal/generator"
	messaging "service-a/internal/messaging/kafka"
	"service-a/internal/router"
	"service-a/internal/service"
	"service-a/internal/validator"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	// Load configs
	appConfig := config.LoadAppConfig()
	sensorGeneratorConfig := config.LoadSensorGeneratorConfig()
	kafkaConfig, err := config.LoadKafkaConfig()
	if err != nil {
		log.Fatalf("❌ Failed to get kafka config: %v", err)
	}

	// Initialize kafka producer
	kafkaProducer := messaging.NewKafkaProducer(kafkaConfig)

	// Init generator with default frequency and start it
	gen := generator.NewGenerator(sensorGeneratorConfig.Frequency, kafkaProducer)

	// Initialize services
	sensorService := service.NewSensorService(kafkaProducer, gen, sensorGeneratorConfig)

	// Initialize controllers
	sensorController := controller.NewSensorController(sensorService)
	ctrl := &controller.Controllers{
		SensorController: sensorController,
	}

	// Create Echo instance
	e := echo.New()
	e.Validator = validator.New()

	// Setup Echo router
	router.NewRouter(e, ctrl)

	// Start server in a goroutine
	go func() {
		log.Printf("🚀 Starting Echo server on port %s", appConfig.AppPort)
		if err := e.Start(":" + appConfig.AppPort); err != nil && err != http.ErrServerClosed {
			log.Printf("❌ Echo server failed: %v", err)
		}
	}()

	// Wait for interrupt signal (graceful shutdown)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down server...")

	// Gracefully shutdown Echo with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Printf("⚠️ Echo forced to shutdown: %v", err)
	}

	// Close kafka producer connection
	if err := kafkaProducer.Close(); err != nil {
		log.Printf("⚠️ Failed to close Kafka broker: %v", err)
	} else {
		log.Println("✅ Kafka broker closed successfully")
	}

	log.Println("👋 Service stopped cleanly")
}

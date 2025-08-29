package config

import (
	"log"
	"os"
	"time"
)

type SensorGeneratorConfig struct {
	Frequency  time.Duration
	Enabled    bool
	SensorType string
}

func LoadSensorGeneratorConfig() *SensorGeneratorConfig {
	raw := os.Getenv("SENSOR_GENERATOR_FREQUENCY")
	if raw == "" {
		log.Println("⚠️ DEFAULT_SENSOR_GENERATOR_FREQUENCY not set, using default 5s")
		raw = "5s"
	}

	freq, err := time.ParseDuration(raw)
	if err != nil {
		log.Printf("⚠️ Invalid DEFAULT_SENSOR_GENERATOR_FREQUENCY=%s, falling back to 5s", raw)
		freq = 5 * time.Second
	}

	sensorType := os.Getenv("SENSOR_TYPE")
	if sensorType == "" {
		sensorType = ""
	}

	return &SensorGeneratorConfig{
		Frequency:  freq,
		Enabled:    os.Getenv("SENSOR_GENERATOR_ENABLED") == "true",
		SensorType: sensorType,
	}
}

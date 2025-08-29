package config

import (
	"os"
	"time"
)

type KafkaConfig struct {
	KafkaBroker   string
	KafkaTopic    string
	KafkaClientID string
	KafkaTimeout  time.Duration
}

func LoadKafkaConfig() (*KafkaConfig, error) {
	timeout, err := time.ParseDuration(os.Getenv("KAFKA_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	return &KafkaConfig{
		KafkaBroker:   os.Getenv("KAFKA_BROKER"),
		KafkaTopic:    os.Getenv("KAFKA_TOPIC"),
		KafkaClientID: os.Getenv("KAFKA_CLIENT_ID"),
		KafkaTimeout:  timeout,
	}, nil
}

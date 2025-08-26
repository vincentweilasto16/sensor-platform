package messaging

import (
	"context"
	"log"
	"time"

	"service-b/internal/config"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(cfg *config.KafkaConfig) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{cfg.KafkaBroker},
			GroupID:  cfg.KafkaGroupID,
			Topic:    cfg.KafkaTopic,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
			MaxWait:  time.Second,
		}),
	}
}

// Consume listens for messages and calls the handler
func (c *KafkaConsumer) Consume(ctx context.Context, handler func(key, message []byte) error) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		log.Printf("📩 Received message: %s", string(msg.Value))

		if handler != nil {
			if err := handler(msg.Key, msg.Value); err != nil {
				log.Printf("❌ Handler error: %v", err)
			}
		}
	}
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}

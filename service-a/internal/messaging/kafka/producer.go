package messaging

import (
	"context"

	"service-a/internal/config"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(cfg *config.KafkaConfig) *KafkaProducer {
	return &KafkaProducer{
		Writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  []string{cfg.KafkaBroker},
			Topic:    cfg.KafkaTopic,
			Balancer: &kafka.LeastBytes{},
		}),
	}
}

func (p *KafkaProducer) Publish(ctx context.Context, key, value []byte) error {
	return p.Writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

func (p *KafkaProducer) Close() error {
	return p.Writer.Close()
}

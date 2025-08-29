package messaging

import (
	"context"
)

//go:generate mockgen -source=./broker.go -destination=./mock/broker_mock.go -package=mock
type Broker interface {
	Consume(ctx context.Context, handler func(key, message []byte) error) error
	Close() error
}

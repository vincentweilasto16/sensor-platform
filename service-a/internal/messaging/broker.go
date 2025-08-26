package messaging

import "context"

//go:generate mockgen -source=./broker.go -destination=./mock/broker_mock.go -package=mock
type Broker interface {
	Publish(ctx context.Context, key, value []byte) error
	Close() error
}

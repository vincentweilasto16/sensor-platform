package service

import (
	"context"
	"encoding/json"
	"service-a/internal/domain"
	"service-a/internal/dto/request"
	"service-a/internal/errors"
	"service-a/internal/messaging"
)

//go:generate mockgen -source=./sensor_service.go -destination=./mock/sensor_service_mock.go -package=mock service-b/internal/service ISensorService
type ISensorService interface {
	GenerateSensor(ctx context.Context, params *request.GenerateSensorRequest) error
}

type SensorService struct {
	Broker messaging.Broker
}

func NewSensorService(broker messaging.Broker) *SensorService {
	return &SensorService{
		Broker: broker,
	}
}

func (s *SensorService) GenerateSensor(ctx context.Context, params *request.GenerateSensorRequest) error {

	msgBytes, err := json.Marshal(domain.SensorData{
		SensorType:   params.SensorType,
		SensorValue:  params.SensorValue,
		DeviceCode:   params.DeviceCode,
		DeviceNumber: params.DeviceNumber,
		Timestamp:    params.Timestamp.UTC(),
	})
	if err != nil {
		return errors.New(errors.InternalServer, "failed to marshal sensor data")
	}

	// Publish to Kafka
	key := params.DeviceCode + string(params.DeviceNumber) // temporary key
	err = s.Broker.Publish(ctx, []byte(key), msgBytes)
	if err != nil {
		return errors.New(errors.InternalServer, "failed to produce message")
	}

	return nil
}

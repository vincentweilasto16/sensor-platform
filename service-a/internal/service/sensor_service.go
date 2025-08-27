package service

import (
	"context"
	"encoding/json"
	"service-a/internal/config"
	"service-a/internal/domain"
	"service-a/internal/dto/request"
	"service-a/internal/errors"
	"service-a/internal/generator"
	"service-a/internal/messaging"
	"strconv"
	"time"
)

//go:generate mockgen -source=./sensor_service.go -destination=./mock/sensor_service_mock.go -package=mock service-b/internal/service ISensorService
type ISensorService interface {
	GenerateSensorManual(ctx context.Context, params *request.GenerateSensorManualRequest) error
	UpdateGenerateSensorFrequency(ctx context.Context, params *request.UpdateGenerateSensorFrequencyRequest) error
	StartSensorGenerator(ctx context.Context) error
	StopSensorGenerator(ctx context.Context) error
}

type SensorService struct {
	Broker                messaging.Broker
	Generator             *generator.Generator
	SensorGeneratorConfig *config.SensorGeneratorConfig
}

func NewSensorService(
	broker messaging.Broker,
	generator *generator.Generator,
	sensorGeneratorConfig *config.SensorGeneratorConfig,
) *SensorService {
	s := &SensorService{
		Broker:                broker,
		Generator:             generator,
		SensorGeneratorConfig: sensorGeneratorConfig,
	}

	if sensorGeneratorConfig.Enabled {
		generator.Start()
	}

	return s
}

func (s *SensorService) GenerateSensorManual(ctx context.Context, params *request.GenerateSensorManualRequest) error {

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
	key := params.DeviceCode + strconv.Itoa(int(params.DeviceNumber)) // temporary key
	err = s.Broker.Publish(ctx, []byte(key), msgBytes)
	if err != nil {
		return errors.New(errors.InternalServer, "failed to produce message")
	}

	return nil
}

func (s *SensorService) UpdateGenerateSensorFrequency(ctx context.Context, params *request.UpdateGenerateSensorFrequencyRequest) error {

	freq := time.Duration(params.Frequency) * time.Second
	s.Generator.UpdateFrequency(freq)

	return nil
}

func (s *SensorService) StartSensorGenerator(ctx context.Context) error {
	s.Generator.Start()
	return nil
}

func (s *SensorService) StopSensorGenerator(ctx context.Context) error {
	s.Generator.Stop()
	return nil
}

package service

import (
	"context"
	"service-b/internal/dto/request"
	"service-b/internal/errors"
	"service-b/internal/repository"
	entity "service-b/internal/repository/mysql"
)

//go:generate mockgen -source=./sensor_service.go -destination=./mock/sensor_service_mock.go -package=mock service-b/internal/service ISensorService
type ISensorService interface {
	GetSensorByDevice(ctx context.Context, params *request.GetSensorByDeviceRequest) ([]*entity.SensorDatum, int64, error)
}

type SensorService struct {
	repo repository.IMySQLRepository
}

func NewSensorService(repo repository.IMySQLRepository) *SensorService {
	return &SensorService{
		repo: repo,
	}
}

func (s *SensorService) GetSensorByDevice(ctx context.Context, params *request.GetSensorByDeviceRequest) ([]*entity.SensorDatum, int64, error) {
	sensorDatas, err := s.repo.GetSensorDataByDeviceCodeAndNumber(ctx, entity.GetSensorDataByDeviceCodeAndNumberParams{
		DeviceCode:   params.DeviceCode,
		DeviceNumber: params.DeviceNumber,
		Limit:        params.Limit,
		Offset:       (params.Page - 1) * params.Limit,
	})
	if err == nil {
		return nil, 0, errors.New(errors.NotFound, "sensors not found")
	}

	total, err := s.repo.CountSensorDataByDeviceCodeAndNumber(ctx, entity.CountSensorDataByDeviceCodeAndNumberParams{
		DeviceCode:   params.DeviceCode,
		DeviceNumber: params.DeviceNumber,
	})
	if err != nil {
		return nil, 0, errors.New(errors.NotFound, "sensors not found")
	}

	sensors := make([]*entity.SensorDatum, len(sensorDatas))
	for i := range sensorDatas {
		sensors[i] = &sensorDatas[i]
	}

	return sensors, total, nil
}

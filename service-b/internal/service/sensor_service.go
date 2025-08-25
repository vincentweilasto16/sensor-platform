package service

import (
	"context"
	"database/sql"
	"service-b/internal/dto/request"
	"service-b/internal/errors"
	"service-b/internal/repository"
	entity "service-b/internal/repository/mysql"
)

//go:generate mockgen -source=./sensor_service.go -destination=./mock/sensor_service_mock.go -package=mock service-b/internal/service ISensorService
type ISensorService interface {
	GetSensors(ctx context.Context, params *request.GetSensorsRequest) ([]*entity.SensorDatum, int64, error)
}

type SensorService struct {
	repo repository.IMySQLRepository
}

func NewSensorService(repo repository.IMySQLRepository) *SensorService {
	return &SensorService{
		repo: repo,
	}
}

func (s *SensorService) GetSensors(ctx context.Context, params *request.GetSensorsRequest) ([]*entity.SensorDatum, int64, error) {

	// validate start time cannot be greater than end time
	if !params.StartTime.IsZero() && !params.EndTime.IsZero() && params.StartTime.After(params.EndTime) {
		return nil, 0, errors.New(errors.BadRequest, "start time cannot be greater than end time")
	}

	total, err := s.repo.CountSensors(ctx, entity.CountSensorsParams{
		DeviceCode: sql.NullString{
			String: params.DeviceCode,
			Valid:  params.DeviceCode != "",
		},
		DeviceNumber: sql.NullInt32{
			Int32: params.DeviceNumber,
			Valid: params.DeviceNumber != 0,
		},
		StartTime: sql.NullTime{
			Time:  params.StartTime,
			Valid: !params.StartTime.IsZero(),
		},
		EndTime: sql.NullTime{
			Time:  params.EndTime,
			Valid: !params.EndTime.IsZero(),
		},
	})
	if err != nil {
		return nil, 0, errors.New(errors.InternalServer, "failed to count sensors")
	}

	// if sensors are empty, early return it
	if total == 0 {
		return []*entity.SensorDatum{}, 0, nil
	}

	sensorDatas, err := s.repo.GetSensors(ctx, entity.GetSensorsParams{
		DeviceCode: sql.NullString{
			String: params.DeviceCode,
			Valid:  params.DeviceCode != "",
		},
		DeviceNumber: sql.NullInt32{
			Int32: params.DeviceNumber,
			Valid: params.DeviceNumber != 0,
		},
		StartTime: sql.NullTime{
			Time:  params.StartTime,
			Valid: !params.StartTime.IsZero(),
		},
		EndTime: sql.NullTime{
			Time:  params.EndTime,
			Valid: !params.EndTime.IsZero(),
		},
		Limit:  params.Limit,
		Offset: (params.Page - 1) * params.Limit,
	})
	if err != nil {
		return nil, 0, errors.New(errors.InternalServer, "failed to fetch sensors")
	}

	sensors := make([]*entity.SensorDatum, len(sensorDatas))
	for i := range sensorDatas {
		sensors[i] = &sensorDatas[i]
	}

	return sensors, total, nil
}

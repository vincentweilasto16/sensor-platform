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
	CreateSensor(ctx context.Context, params *request.CreateSensorRequest) error
	DeleteSensors(ctx context.Context, params *request.DeleteSensorsRequest) error
	UpdateSensors(ctx context.Context, params *request.UpdateSensorsRequest) error
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

func (s *SensorService) CreateSensor(ctx context.Context, params *request.CreateSensorRequest) error {
	err := s.repo.InsertSensorData(ctx, entity.InsertSensorDataParams{
		SensorType:   params.SensorType,
		SensorValue:  params.SensorValue,
		DeviceCode:   params.DeviceCode,
		DeviceNumber: params.DeviceNumber,
		Timestamp:    params.Timestamp,
	})
	if err != nil {
		return errors.New(errors.InternalServer, "failed to store sensor")
	}

	return nil
}

func (s *SensorService) DeleteSensors(ctx context.Context, params *request.DeleteSensorsRequest) error {
	if params.DeviceCode == "" && params.DeviceNumber == 0 && params.StartTime.IsZero() && params.EndTime.IsZero() {
		return errors.New(errors.BadRequest, "at least one filter is required for deletion")
	}

	// validate start time cannot be greater than end time
	if !params.StartTime.IsZero() && !params.EndTime.IsZero() && params.StartTime.After(params.EndTime) {
		return errors.New(errors.BadRequest, "start time cannot be greater than end time")
	}

	err := s.repo.DeleteSensors(ctx, entity.DeleteSensorsParams{
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
		return errors.New(errors.InternalServer, "failed to delete sensors")
	}

	return nil
}

func (s *SensorService) UpdateSensors(ctx context.Context, params *request.UpdateSensorsRequest) error {
	// Validate at least one filter is provided
	if params.Criteria.DeviceCode == "" &&
		params.Criteria.DeviceNumber == 0 &&
		params.Criteria.StartTime.IsZero() &&
		params.Criteria.EndTime.IsZero() {
		return errors.New(errors.BadRequest, "at least one filter is required for update")
	}

	// validate start time cannot be greater than end time
	if !params.Criteria.StartTime.IsZero() && !params.Criteria.EndTime.IsZero() && params.Criteria.StartTime.After(params.Criteria.EndTime) {
		return errors.New(errors.BadRequest, "start time cnnaot be greater than end time")
	}

	// Validate at least one change is provided
	if params.Changes.SensorValue == 0 &&
		params.Changes.SensorType == "" &&
		params.Changes.Timestamp.IsZero() {
		return errors.New(errors.BadRequest, "at least one field to update must be provided")
	}

	err := s.repo.UpdateSensors(ctx, entity.UpdateSensorsParams{
		SensorValue: sql.NullFloat64{
			Float64: params.Changes.SensorValue,
			Valid:   params.Changes.SensorValue != 0,
		},
		SensorType: sql.NullString{
			String: params.Changes.SensorType,
			Valid:  params.Changes.SensorType != "",
		},
		Timestamp: sql.NullTime{
			Time:  params.Changes.Timestamp,
			Valid: !params.Changes.Timestamp.IsZero(),
		},
		DeviceCode: sql.NullString{
			String: params.Criteria.DeviceCode,
			Valid:  params.Criteria.DeviceCode != "",
		},
		DeviceNumber: sql.NullInt32{
			Int32: params.Criteria.DeviceNumber,
			Valid: params.Criteria.DeviceNumber != 0,
		},
		StartTime: sql.NullTime{
			Time:  params.Criteria.StartTime,
			Valid: !params.Criteria.StartTime.IsZero(),
		},
		EndTime: sql.NullTime{
			Time:  params.Criteria.EndTime,
			Valid: !params.Criteria.EndTime.IsZero(),
		},
	})
	if err != nil {
		return errors.New(errors.InternalServer, "failed to update sensors")
	}

	return nil
}

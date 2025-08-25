package repository

import (
	"context"
	entity "service-b/internal/repository/mysql"
)

//go:generate mockgen -source=./mysql_repository.go -destination=./mock/mysql_repository_mock.go -package=mock
type IMySQLRepository interface {

	// Sensor Data Repository
	GetSensors(ctx context.Context, arg entity.GetSensorsParams) ([]entity.SensorDatum, error)
	CountSensors(ctx context.Context, arg entity.CountSensorsParams) (int64, error)
	InsertSensorData(ctx context.Context, arg entity.InsertSensorDataParams) error
	UpdateSensorData(ctx context.Context, arg entity.UpdateSensorDataParams) error
	DeleteSensorData(ctx context.Context, arg entity.DeleteSensorDataParams) error
}

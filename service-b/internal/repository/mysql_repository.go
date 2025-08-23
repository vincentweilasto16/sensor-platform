package repository

import (
	"context"
	entity "service-b/internal/repository/mysql"
)

//go:generate mockgen -source=./mysql_repository.go -destination=./mock/mysql_repository_mock.go -package=mock
type IMySQLRepository interface {

	// Sensor Data Repository
	GetSensorDataByTime(ctx context.Context, arg entity.GetSensorDataByTimeParams) ([]entity.SensorDatum, error)
	GetSensorDataByDeviceCodeAndNumber(ctx context.Context, arg entity.GetSensorDataByDeviceCodeAndNumberParams) ([]entity.SensorDatum, error)
	GetSensorDataByDeviceAndTime(ctx context.Context, arg entity.GetSensorDataByDeviceAndTimeParams) ([]entity.SensorDatum, error)
	CountSensorDataByDeviceCodeAndNumber(ctx context.Context, arg entity.CountSensorDataByDeviceCodeAndNumberParams) (int64, error)
	InsertSensorData(ctx context.Context, arg entity.InsertSensorDataParams) error
	UpdateSensorData(ctx context.Context, arg entity.UpdateSensorDataParams) error
	DeleteSensorData(ctx context.Context, arg entity.DeleteSensorDataParams) error
}

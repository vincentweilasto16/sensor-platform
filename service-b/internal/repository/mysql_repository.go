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
	UpdateSensors(ctx context.Context, arg entity.UpdateSensorsParams) error
	DeleteSensors(ctx context.Context, arg entity.DeleteSensorsParams) error

	// User Repository
	InsertUser(ctx context.Context, arg entity.InsertUserParams) error
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}

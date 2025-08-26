package request

import (
	"time"
)

type GenerateSensorRequest struct {
	SensorValue  float64   `json:"sensor_value" validate:"required"`
	SensorType   string    `json:"sensor_type" validate:"required"`
	DeviceCode   string    `json:"device_code" validate:"required"`
	DeviceNumber int32     `json:"device_number" validate:"required"`
	Timestamp    time.Time `json:"timestamp" validate:"required"`
}

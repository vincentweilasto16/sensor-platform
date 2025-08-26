package request

import (
	"time"
)

type GetSensorsRequest struct {
	DeviceCode   string    `json:"device_code" query:"device_code" validate:"omitempty"`
	DeviceNumber int32     `json:"device_number" query:"device_number" validate:"omitempty"`
	StartTime    time.Time `json:"start_time" query:"start_time" validate:"omitempty"`
	EndTime      time.Time `json:"end_time" query:"end_time" validate:"omitempty"`
	Limit        int32     `json:"limit" query:"limit" validate:"required,min=1"`
	Page         int32     `json:"page" query:"page" validate:"required,min=1"`
}

type CreateSensorRequest struct {
	SensorValue  float64   `json:"sensor_value" validate:"required"`
	SensorType   string    `json:"sensor_type" validate:"required"`
	DeviceCode   string    `json:"device_code" validate:"required"`
	DeviceNumber int32     `json:"device_number" validate:"required"`
	Timestamp    time.Time `json:"timestamp" validate:"required"`
}

type UpdateSensorsCriteriaRequest struct {
	DeviceCode   string    `json:"device_code,omitempty"`
	DeviceNumber int32     `json:"device_number,omitempty"`
	StartTime    time.Time `json:"start_time,omitempty"`
	EndTime      time.Time `json:"end_time,omitempty"`
}

type UpdateSensorsChangesRequest struct {
	SensorValue float64   `json:"sensor_value" validate:"omitempty"`
	SensorType  string    `json:"sensor_type" validate:"omitempty"`
	Timestamp   time.Time `json:"timestamp" validate:"omitempty"`
}

type UpdateSensorsRequest struct {
	Criteria UpdateSensorsCriteriaRequest `json:"criteria" validate:"required"`
	Changes  UpdateSensorsChangesRequest  `json:"changes" validate:"required"`
}

type DeleteSensorsRequest struct {
	DeviceCode   string    `json:"device_code" validate:"omitempty"`
	DeviceNumber int32     `json:"device_number" validate:"omitempty"`
	StartTime    time.Time `json:"start_time" validate:"omitempty"`
	EndTime      time.Time `json:"end_time" query:"end_time" validate:"omitempty"`
}

package response

import (
	"github.com/guregu/null"
)

type SensorResponse struct {
	ID           int64       `json:"id"`
	SensorType   string      `json:"sensor_type"`
	SensorValue  float64     `json:"sensor_value"`
	DeviceCode   string      `json:"device_code"`
	DeviceNumber int32       `json:"device_number"`
	Timestamp    string      `json:"timestamp"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	DeletedAt    null.String `json:"deleted_at"`
}

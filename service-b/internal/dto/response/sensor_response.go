package response

import (
	"time"

	"github.com/guregu/null"
)

type SensorResponse struct {
	ID           int64     `json:"id"`
	SensorType   string    `json:"sensor_type"`
	SensorValue  float64   `json:"sensor_value"`
	DeviceCode   string    `json:"device_code"`
	DeviceNumber int32     `json:"device_number"`
	Timestamp    time.Time `json:"timestamp"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    null.Time `json:"deleted_at"`
}

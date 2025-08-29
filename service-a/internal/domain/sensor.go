package domain

import "time"

type SensorData struct {
	SensorType   string    `json:"sensor_type"`
	SensorValue  float64   `json:"sensor_value"`
	DeviceCode   string    `json:"device_code"`
	DeviceNumber int32     `json:"device_number"`
	Timestamp    time.Time `json:"timestamp"`
}

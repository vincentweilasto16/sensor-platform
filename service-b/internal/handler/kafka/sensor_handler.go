package kafka

import (
	"context"
	"encoding/json"
	"log"

	"service-b/internal/dto/request"
	"service-b/internal/service"
)

type SensorHandler struct {
	SensorService *service.SensorService
}

func NewSensorHandler(sensorService *service.SensorService) *SensorHandler {
	return &SensorHandler{
		SensorService: sensorService,
	}
}

func (h *SensorHandler) ProcessSensorEvent(ctx context.Context, msg []byte) error {
	var payload request.CreateSensorRequest
	if err := json.Unmarshal(msg, &payload); err != nil {
		log.Printf("❌ Failed to unmarshal message: %v", err)
		return err
	}

	err := h.SensorService.CreateSensor(ctx, &payload)
	if err != nil {
		log.Printf("❌ Failed to store sensor data: %v", err)
		return err
	}

	log.Printf("✅ Successfully processed message for device %s-%d", payload.DeviceCode, payload.DeviceNumber)
	return nil
}

package presenter

import (
	"service-b/internal/dto/response"
	entity "service-b/internal/repository/mysql"
	"time"

	"github.com/guregu/null"
)

func SensorResponse(e *entity.SensorDatum) *response.SensorResponse {
	if e == nil {
		return nil
	}

	deletedAt := null.NewString("", false)
	if e.DeletedAt.Valid {
		deletedAt = null.NewString(e.DeletedAt.Time.UTC().Format(time.RFC3339), true)
	}

	return &response.SensorResponse{
		ID:           e.ID,
		SensorType:   e.SensorType,
		SensorValue:  e.SensorValue,
		DeviceCode:   e.DeviceCode,
		DeviceNumber: e.DeviceNumber,
		Timestamp:    e.Timestamp.UTC().Format(time.RFC3339),
		CreatedAt:    e.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:    e.UpdatedAt.UTC().Format(time.RFC3339),
		DeletedAt:    deletedAt,
	}
}

func SensorsResponse(datas []*entity.SensorDatum) []*response.SensorResponse {
	res := []*response.SensorResponse{}
	for _, e := range datas {
		res = append(res, SensorResponse(e))
	}
	return res
}

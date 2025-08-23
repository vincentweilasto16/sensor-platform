package presenter

import (
	"service-b/internal/dto/response"
	entity "service-b/internal/repository/mysql"

	"github.com/guregu/null"
)

func SensorResponse(e *entity.SensorDatum) *response.SensorResponse {
	if e == nil {
		return nil
	}

	return &response.SensorResponse{
		ID:           e.ID,
		SensorType:   e.SensorType,
		SensorValue:  e.SensorValue,
		DeviceCode:   e.DeviceCode,
		DeviceNumber: e.DeviceNumber,
		Timestamp:    e.Timestamp,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
		DeletedAt:    null.NewTime(e.DeletedAt.Time, e.DeletedAt.Valid),
	}
}

func SensorsResponse(datas []*entity.SensorDatum) []*response.SensorResponse {
	res := []*response.SensorResponse{}
	for _, e := range datas {
		res = append(res, SensorResponse(e))
	}
	return res
}

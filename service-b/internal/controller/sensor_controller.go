package controller

import (
	"service-b/internal/dto/request"
	"service-b/internal/dto/response"
	"service-b/internal/pagination"
	"service-b/internal/presenter"
	"service-b/internal/service"

	"github.com/labstack/echo/v4"
)

type SensorController struct {
	SensorService service.ISensorService
}

func NewSensorController(SensorService service.ISensorService) *SensorController {
	return &SensorController{
		SensorService: SensorService,
	}
}

func (c *SensorController) GetSensors(ctx echo.Context) error {
	// @TODO: prepare the context

	var queryParams request.GetSensorsRequest
	if err := request.SetQueryParams(ctx, &queryParams); err != nil {
		return response.RespondError(ctx, err)
	}

	if err := ctx.Validate(&queryParams); err != nil {
		return response.RespondError(ctx, err)
	}

	sensorData, total, err := c.SensorService.GetSensors(ctx.Request().Context(), &queryParams)
	if err != nil {
		return response.RespondError(ctx, err)
	}

	pagination := pagination.BuildPaginator(int(total), int(queryParams.Limit), int((queryParams.Page-1))*int(queryParams.Limit))
	res := presenter.SensorsResponse(sensorData)

	return response.RespondSuccessWithPaginator(ctx, res, pagination, "ok")
}

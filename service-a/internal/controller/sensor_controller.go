package controller

import (
	"service-a/internal/dto/request"
	"service-a/internal/dto/response"

	"service-a/internal/service"

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

func (c *SensorController) GenerateSensor(ctx echo.Context) error {
	// @TODO: prepare the context

	var bodyRequest request.GenerateSensorRequest
	if err := request.SetBodyParams(ctx, &bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	if err := ctx.Validate(&bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	err := c.SensorService.GenerateSensor(ctx.Request().Context(), &bodyRequest)
	if err != nil {
		return response.RespondError(ctx, err)
	}

	return response.RespondSuccess(ctx, nil, "ok")
}

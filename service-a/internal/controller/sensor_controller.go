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

func (c *SensorController) GenerateSensorManual(ctx echo.Context) error {
	// @TODO: prepare the context

	var bodyRequest request.GenerateSensorManualRequest
	if err := request.SetBodyParams(ctx, &bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	if err := ctx.Validate(&bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	err := c.SensorService.GenerateSensorManual(ctx.Request().Context(), &bodyRequest)
	if err != nil {
		return response.RespondError(ctx, err)
	}

	return response.RespondSuccess(ctx, nil, "ok")
}

func (c *SensorController) UpdateGenerateSensorFrequency(ctx echo.Context) error {
	// @TODO: prepare the context

	var bodyRequest request.UpdateGenerateSensorFrequencyRequest
	if err := request.SetBodyParams(ctx, &bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	if err := ctx.Validate(&bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	err := c.SensorService.UpdateGenerateSensorFrequency(ctx.Request().Context(), &bodyRequest)
	if err != nil {
		return response.RespondError(ctx, err)
	}

	return response.RespondSuccess(ctx, nil, "ok")
}

func (c *SensorController) StartSensorGenerator(ctx echo.Context) error {
	if err := c.SensorService.StartSensorGenerator(ctx.Request().Context()); err != nil {
		return response.RespondError(ctx, err)
	}
	return response.RespondSuccess(ctx, nil, "sensor generator started")
}

func (c *SensorController) StopSensorGenerator(ctx echo.Context) error {
	if err := c.SensorService.StopSensorGenerator(ctx.Request().Context()); err != nil {
		return response.RespondError(ctx, err)
	}
	return response.RespondSuccess(ctx, nil, "sensor generator stopped")
}


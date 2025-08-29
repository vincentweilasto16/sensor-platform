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

func (c *SensorController) DeleteSensors(ctx echo.Context) error {
	// @TODO: prepare the context

	var bodyRequest request.DeleteSensorsRequest
	if err := request.SetBodyParams(ctx, &bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	if err := ctx.Validate(&bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	err := c.SensorService.DeleteSensors(ctx.Request().Context(), &bodyRequest)
	if err != nil {
		return response.RespondError(ctx, err)
	}

	return response.RespondSuccess(ctx, nil, "ok")
}

func (c *SensorController) UpdateSensors(ctx echo.Context) error {
	// @TODO: prepare the context

	var bodyRequest request.UpdateSensorsRequest
	if err := request.SetBodyParams(ctx, &bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	if err := ctx.Validate(&bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	err := c.SensorService.UpdateSensors(ctx.Request().Context(), &bodyRequest)
	if err != nil {
		return response.RespondError(ctx, err)
	}

	return response.RespondSuccess(ctx, nil, "ok")
}

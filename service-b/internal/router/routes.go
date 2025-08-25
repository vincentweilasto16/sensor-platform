package router

import (
	"service-b/internal/constants"
	"service-b/internal/controller"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(rg *echo.Group, ctrl *controller.Controllers) {

	// sensor route
	sensors := rg.Group(constants.SensorBasePath)
	{
		sensors.GET("", ctrl.SensorController.GetSensors)
	}
}

package router

import (
	"service-a/internal/constants"
	"service-a/internal/controller"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(rg *echo.Group, ctrl *controller.Controllers) {

	// sensor route
	sensors := rg.Group(constants.SensorBasePath)
	{
		sensors.POST("/generate/manual", ctrl.SensorController.GenerateSensorManual)
		sensors.PUT("/generate/frequency", ctrl.SensorController.UpdateGenerateSensorFrequency)
		sensors.POST("/generate/start", ctrl.SensorController.StartSensorGenerator)
		sensors.POST("/generate/stop", ctrl.SensorController.StopSensorGenerator)
	}
}

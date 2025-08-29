package router

import (
	"service-b/internal/config"
	"service-b/internal/constants"
	"service-b/internal/controller"
	"service-b/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(rg *echo.Group, ctrl *controller.Controllers, jwtConfig *config.JWTConfig) {
	// sensor route
	sensors := rg.Group(constants.SensorBasePath)
	sensors.Use(middleware.JWTMiddleware(jwtConfig))
	{
		sensors.GET("", ctrl.SensorController.GetSensors)
		sensors.PUT("", ctrl.SensorController.UpdateSensors, middleware.RolesAllowed("admin", "user"))
		sensors.DELETE("", ctrl.SensorController.DeleteSensors, middleware.RolesAllowed("admin"))
	}

	// auth route (public)
	auth := rg.Group(constants.AuthBasePath)
	{
		auth.POST("/register", ctrl.AuthController.Register)
		auth.POST("/login", ctrl.AuthController.Login)
	}
}

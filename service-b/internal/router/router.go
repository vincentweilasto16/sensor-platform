package router

import (
	"service-b/internal/config"
	"service-b/internal/constants"
	"service-b/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, ctrl *controller.Controllers, jwtConfig *config.JWTConfig) {
	publicV1 := e.Group(constants.PublicAPIV1BasePath)
	RegisterRoutes(publicV1, ctrl, jwtConfig)
}

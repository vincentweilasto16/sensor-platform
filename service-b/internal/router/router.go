package router

import (
	"service-b/internal/constants"
	"service-b/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, ctrl *controller.Controllers) {
	// Public API v1
	publicV1 := e.Group(constants.PublicAPIV1BasePath)
	RegisterRoutes(publicV1, ctrl)
}

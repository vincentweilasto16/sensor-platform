package response

import (
	"net/http"
	"service-a/internal/errors"

	"github.com/labstack/echo/v4"
)

// Standard API response
type StandardResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error response
type ErrorResponse struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Fields  map[string][]string `json:"fields,omitempty"`
}

// Send success response
func RespondSuccess(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusOK, StandardResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Send error response
func RespondError(c echo.Context, err error) error {
	appErr, ok := err.(errors.AppError)
	status := http.StatusInternalServerError
	if ok {
		status = int(appErr.Type)
	}

	return c.JSON(status, ErrorResponse{
		Status:  "error",
		Message: err.Error(),
		Fields:  errors.GetFields(err),
	})
}

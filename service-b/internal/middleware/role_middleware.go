package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RolesAllowed(allowedRoles ...string) echo.MiddlewareFunc {
	roleMap := make(map[string]struct{})
	for _, r := range allowedRoles {
		roleMap[r] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := GetUserRole(c)
			if role == "" {
				return echo.NewHTTPError(http.StatusForbidden, "role not found")
			}

			if _, exists := roleMap[role]; !exists {
				return echo.NewHTTPError(http.StatusForbidden, "access denied for this role")
			}

			return next(c)
		}
	}
}

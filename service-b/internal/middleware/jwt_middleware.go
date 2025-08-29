package middleware

import (
	"errors"
	"net/http"
	"service-b/internal/config"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func JWTMiddleware(jwtConfig *config.JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header")
			}

			tokenStr := parts[1]
			token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtConfig.Secret), nil
			})
			if err != nil {
				// check if it's specifically an expiration error
				if errors.Is(err, jwt.ErrTokenExpired) {
					return echo.NewHTTPError(http.StatusUnauthorized, "token expired")
				}
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			if !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			claims, ok := token.Claims.(*JWTClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
			}

			// store user info in context
			c.Set("user_id", claims.UserID)
			c.Set("role", claims.Role)

			return next(c)
		}
	}
}

func GetUserID(c echo.Context) int64 {
	if val, ok := c.Get("user_id").(int64); ok {
		return val
	}
	return 0
}

func GetUserRole(c echo.Context) string {
	if val, ok := c.Get("role").(string); ok {
		return val
	}
	return ""
}

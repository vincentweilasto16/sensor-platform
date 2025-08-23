package request

import (
	"github.com/labstack/echo/v4"
)

func SetBodyParams(c echo.Context, target interface{}) error {
	if err := c.Bind(target); err != nil {
		return err
	}
	return nil
}

func SetQueryParams(c echo.Context, target interface{}) error {
	if err := c.Bind(target); err != nil {
		return err
	}
	return nil
}

func SetURIParams(c echo.Context, target interface{}) error {
	if err := c.Bind(target); err != nil {
		return err
	}
	return nil
}

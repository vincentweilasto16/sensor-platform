package controller

import (
	"service-b/internal/dto/request"
	"service-b/internal/dto/response"
	"service-b/internal/service"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService service.IAuthService
}

func NewAuthController(AuthService service.IAuthService) *AuthController {
	return &AuthController{
		AuthService: AuthService,
	}
}

func (c *AuthController) Register(ctx echo.Context) error {
	// @TODO: prepare the context

	var bodyRequest request.RegisterRequest
	if err := request.SetBodyParams(ctx, &bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	if err := ctx.Validate(&bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	err := c.AuthService.Register(ctx.Request().Context(), &bodyRequest)
	if err != nil {
		return response.RespondError(ctx, err)
	}

	return response.RespondSuccess(ctx, nil, "ok")
}

func (c *AuthController) Login(ctx echo.Context) error {
	// @TODO: prepare the context

	var bodyRequest request.LoginRequest
	if err := request.SetBodyParams(ctx, &bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	if err := ctx.Validate(&bodyRequest); err != nil {
		return response.RespondError(ctx, err)
	}

	res, err := c.AuthService.Login(ctx.Request().Context(), &bodyRequest)
	if err != nil {
		return response.RespondError(ctx, err)
	}

	return response.RespondSuccess(ctx, res, "ok")
}

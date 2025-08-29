package service

import (
	"context"
	"service-b/internal/config"
	"service-b/internal/dto/request"
	"service-b/internal/dto/response"
	"service-b/internal/errors"
	"service-b/internal/repository"
	entity "service-b/internal/repository/mysql"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=./auth_service.go -destination=./mock/auth_service_mock.go -package=mock service-b/internal/service IAuthService
type IAuthService interface {
	Register(ctx context.Context, params *request.RegisterRequest) error
	Login(ctx context.Context, params *request.LoginRequest) (*response.LoginResponse, error)
}

type AuthService struct {
	repo      repository.IMySQLRepository
	jwtConfig *config.JWTConfig
}

func NewAuthService(repo repository.IMySQLRepository, jwtConfig *config.JWTConfig) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtConfig: jwtConfig,
	}
}

func (s *AuthService) Register(ctx context.Context, params *request.RegisterRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(errors.InternalServer, "failed to hash password")
	}

	err = s.repo.InsertUser(ctx, entity.InsertUserParams{
		Username: params.Username,
		Password: string(hash),
		Role:     entity.UsersRole(params.Role),
	})
	if err != nil {
		return errors.New(errors.InternalServer, "failed to insert user")
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, params *request.LoginRequest) (*response.LoginResponse, error) {
	user, err := s.repo.GetUserByUsername(ctx, params.Username)
	if err != nil {
		return nil, errors.New(errors.NotFound, "user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return nil, errors.New(errors.BadRequest, "invalid password")
	}

	exp := time.Now().Add(time.Duration(s.jwtConfig.ExpiresIn) * time.Second).Unix()

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return nil, errors.New(errors.InternalServer, "failed to sign JWT token")
	}

	return &response.LoginResponse{
		AccessToken: tokenString,
		ExpiresAt:   exp,
	}, nil
}

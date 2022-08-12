package ports

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/beta-team/internal/requests"
)

type IUserService interface {
	//Login(ctx context.Context, input *requests.LoginRequest) (*string, error)
	Register(ctx context.Context, input *requests.RegisterRequest) error
}

type IUserRepository interface {
	//Login(ctx context.Context, input *requests.LoginRequest) (string, error)
	CreateUser(ctx context.Context, input *requests.CreateUserRequest) error
}

type IUserHandlers interface {
	Login(c echo.Context) error
	Register(c *echo.Context) error
}

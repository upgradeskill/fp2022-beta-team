package user

import (
	"github.com/labstack/echo/v4"
	userPort "github.com/upgradeskill/beta-team/internal/core/ports/user"
)

type UserHandler struct {
	userService userPort.IUserService
}

func NewUserHandler(userService userPort.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Login(ctx echo.Context) error {

	return nil
}

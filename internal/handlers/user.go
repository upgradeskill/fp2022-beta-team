package user

import (
	"github.com/labstack/echo/v4"
	userPort "github.com/upgradeskill/beta-team/internal/core/ports"
)

type UserHandler struct {
	userUsecase userPort.IUserUsecase
}

func NewUserHandler(userUsecase userPort.IUserUsecase) UserHandler {
	return UserHandler{
		userUsecase: userUsecase,
	}
}

func (uh *UserHandler) Register(ctx echo.Context) error {

	return nil
}

func (uh *UserHandler) Login(ctx echo.Context) error {

	return nil
}

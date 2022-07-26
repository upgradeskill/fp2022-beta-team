package user

import (
	userPort "github.com/upgradeskill/beta-team/internal/core/ports/user"
)

type UserHandler struct {
	userService userPort.UserService
}

func NewUserHandler(userService userPort.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

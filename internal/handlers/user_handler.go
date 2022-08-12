package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/beta-team/conf"
	port "github.com/upgradeskill/beta-team/internal/core/ports"
	"github.com/upgradeskill/beta-team/internal/requests"
	"github.com/upgradeskill/beta-team/pkg/middleware"
	"github.com/upgradeskill/beta-team/pkg/mlog"
	"github.com/upgradeskill/beta-team/pkg/mvalidator"
	"github.com/upgradeskill/beta-team/pkg/web"
)

type UserHandler struct {
	userService port.IUserService
	Validator   mvalidator.Validator
	Logger      mlog.Logger
	Cfg         conf.Group
}

func NewUserHandler(
	userService port.IUserService,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config conf.Group,
) *UserHandler {
	return &UserHandler{
		userService: userService,
		Validator:   validator,
		Logger:      logger,
		Cfg:         config,
	}
}

func (uh *UserHandler) Register(ctx echo.Context) error {
	requestID := middleware.GetID(ctx)
	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

	var payloads requests.RegisterRequest
	if err := ctx.Bind(&payloads); err != nil {
		uh.Logger.ErrorT(requestID, "register payload", err, mlog.Any("payload", payloads))
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	mapErr, err := uh.Validator.SliceStruct(payloads)
	if err != nil {
		uh.Logger.ErrorT(requestID, "Bad Request", err)
		return web.ResponseErrValidation(ctx, "bad request", mapErr)
	}

	err = uh.userService.Register(userCtx, &payloads)
	if err != nil {
		uh.Logger.ErrorT(requestID, "error register user", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

// func (uh *UserHandler) Login(ctx echo.Context) error {
// 	requestID := middleware.GetID(ctx)
// 	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

// 	var payloads requests.LoginRequest
// 	if err := ctx.Bind(&payloads); err != nil {
// 		uh.Logger.ErrorT(requestID, "login payload", err, mlog.Any("payload", payloads))
// 		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
// 	}

// 	mapErr, err := uh.Validator.SliceStruct(payloads)
// 	if err != nil {
// 		uh.Logger.ErrorT(requestID, "Bad Request", err)
// 		return web.ResponseErrValidation(ctx, "bad request", mapErr)
// 	}

// 	result, err := uh.userService.Login(userCtx, &payloads)
// 	if err != nil {
// 		uh.Logger.ErrorT(requestID, "error register user", err)
// 		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
// 	}

// 	return web.ResponseFormatter(ctx, http.StatusOK, "Success", result, nil)
// }

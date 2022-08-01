package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/beta-team/conf"
	userPort "github.com/upgradeskill/beta-team/internal/core/ports"
	"github.com/upgradeskill/beta-team/internal/requests"
	"github.com/upgradeskill/beta-team/pkg/middleware"
	"github.com/upgradeskill/beta-team/pkg/mlog"
	"github.com/upgradeskill/beta-team/pkg/mvalidator"
	"github.com/upgradeskill/beta-team/pkg/web"
)

type UserHandler struct {
	userUseCase userPort.IUserUsecase
	Validator   mvalidator.Validator
	Logger      mlog.Logger
	Cfg         conf.Group
}

func NewUserHandler(
	userUsecase userPort.IUserUsecase,
	validator mvalidator.Validator,
	logger mlog.Logger,
	config conf.Group,
) *UserHandler {
	return &UserHandler{
		userUseCase: userUsecase,
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

	err = uh.userUseCase.RegisterUser(userCtx, &payloads)
	if err != nil {
		uh.Logger.ErrorT(requestID, "error register user", err)
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

func (uh *UserHandler) Login(ctx echo.Context) error {

	return nil
}

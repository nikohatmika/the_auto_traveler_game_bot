package player_auth

import (
	"auto_traveler/bussiness/player_auth"
	"auto_traveler/controller"
	playerAuthReq "auto_traveler/controller/player_auth/request"
	"auto_traveler/controller/player_auth/response"
	playerReq "auto_traveler/controller/players/request"

	"net/http"

	"github.com/labstack/echo/v4"
)

type PlayerAuthController struct {
	playerAuthUsecase player_auth.Usecase
}

func NewplayerAuthController(e *echo.Echo, uc player_auth.Usecase) *PlayerAuthController {
	return &PlayerAuthController{
		playerAuthUsecase: uc,
	}
}

func (ctrl *PlayerAuthController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := playerReq.Player{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.playerAuthUsecase.Register(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnauthorized, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&resp))

}

func (ctrl *PlayerAuthController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := playerAuthReq.PlayerAuth{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.playerAuthUsecase.Login(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnauthorized, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&resp))

}

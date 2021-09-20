package players

import (
	"auto_traveler/app/middleware"
	"auto_traveler/bussiness/players"
	"auto_traveler/controller"
	"auto_traveler/controller/players/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlayersController struct {
	playerUsecase players.Usecase
}

func NewPlayerController(e *echo.Echo, uc players.Usecase) *PlayersController {
	return &PlayersController{
		playerUsecase: uc,
	}
}

func (ctrl *PlayersController) FindByToken(c echo.Context) error {
	fmt.Println("ISI C", c)

	ctx := c.Request().Context()
	player := middleware.GetPlayer(c)

	res, err := ctrl.playerUsecase.FindByID(ctx, player.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&res))
}

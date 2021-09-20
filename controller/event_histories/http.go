package event_histories

import (
	"auto_traveler/app/middleware"
	"auto_traveler/bussiness/event_histories"
	"auto_traveler/controller"
	"auto_traveler/controller/event_histories/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventHistoriesController struct {
	eventHistoriesUsecase event_histories.Usecase
}

func NewEventHistoriesController(e *echo.Echo, u event_histories.Usecase) *EventHistoriesController {
	return &EventHistoriesController{
		eventHistoriesUsecase: u,
	}
}

func (ctrl *EventHistoriesController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := ctrl.eventHistoriesUsecase.Find(ctx, 0)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	res := []response.EventHistories{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, res)
}

func (ctrl *EventHistoriesController) FindByToken(c echo.Context) error {
	ctx := c.Request().Context()
	player := middleware.GetPlayer(c)
	fmt.Println("PLAYER ID", player.ID)

	data, _ := ctrl.eventHistoriesUsecase.Find(ctx, player.ID)
	// if err != nil {
	// 	return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	// }
	fmt.Println("DATA", data)


	res := []response.EventHistories{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}
	fmt.Println("RES", res)


	return controller.NewSuccessResponse(c, res)
}
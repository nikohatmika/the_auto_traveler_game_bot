package events

import (
	"auto_traveler/bussiness/events"
	"auto_traveler/controller"
	"auto_traveler/controller/events/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventsController struct {
	eventsUsecase events.Usecase
}

func NewEventsController(e *echo.Echo, u events.Usecase) *EventsController {
	return &EventsController{
		eventsUsecase: u,
	}
}

func (ctrl *EventsController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	eventType := c.QueryParam("type")

	data, err := ctrl.eventsUsecase.Find(ctx, eventType)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	res := []response.Events{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, res)
}
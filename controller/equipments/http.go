package equipments

import (
	"auto_traveler/bussiness/equipments"
	"auto_traveler/controller"
	"auto_traveler/controller/equipments/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EquipmentsController struct {
	equipmentsUsecase equipments.Usecase
}

func NewEquipmentsController(e *echo.Echo, u equipments.Usecase) *EquipmentsController {
	return &EquipmentsController{
		equipmentsUsecase: u,
	}
}

func (ctrl *EquipmentsController) Find(c echo.Context) error {
	ctx := c.Request().Context()

	eqType := c.QueryParam("type")

	data, err := ctrl.equipmentsUsecase.Find(ctx, eqType)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	res := []response.Equipments{}
	for _, value := range data {
		res = append(res, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, res)
}
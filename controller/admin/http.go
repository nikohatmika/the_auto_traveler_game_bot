package admin

import (
	"auto_traveler/app/middleware"
	"auto_traveler/bussiness/admin"
	"auto_traveler/controller"
	"auto_traveler/controller/admin/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUsecase admin.Usecase
}

func NewAdminController(e *echo.Echo, u admin.Usecase) *AdminController {
	return &AdminController{
		adminUsecase: u,
	}
}

func (ctrl *AdminController) FindByToken(c echo.Context) error {
	ctx := c.Request().Context()
	admin := middleware.GetAdmin(c)

	res, err := ctrl.adminUsecase.FindByID(ctx, admin.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&res))
}

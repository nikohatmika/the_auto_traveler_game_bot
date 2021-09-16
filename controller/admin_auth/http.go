package admin_auth

import (
	"auto_traveler/bussiness/admin_auth"
	"auto_traveler/controller"
	"auto_traveler/controller/admin_auth/request"
	"auto_traveler/controller/admin_auth/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminAuthController struct {
	adminAuthUsecase admin_auth.Usecase
}

func NewAdminAuthController(e *echo.Echo, cu admin_auth.Usecase) *AdminAuthController {
	return &AdminAuthController{
		adminAuthUsecase: cu,
	}
}

func (ctrl *AdminAuthController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.AdminAuth{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.adminAuthUsecase.Login(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnauthorized, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&resp))

}

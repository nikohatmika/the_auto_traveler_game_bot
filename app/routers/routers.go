package routers

import (
	"auto_traveler/app/middleware"
	"auto_traveler/controller/admin"
	"auto_traveler/controller/admin_auth"
	"auto_traveler/controller/player_auth"
	"auto_traveler/controller/players"

	"auto_traveler/controller/equipments"
	"auto_traveler/controller/events"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	JWTMiddleware           *middleware.ConfigJWT
	AdminAuthController     admin_auth.AdminAuthController
	PlayerAuthController  	player_auth.PlayerAuthController
	AdminController     	admin.AdminController
	PlayersController     	players.PlayersController
	EventsController 		events.EventsController
	EquipmentsController 	equipments.EquipmentsController
}

func (c *ControllerList) RouteRegister(e *echo.Echo) {
	// API USER
	playerMiddleware := *c.JWTMiddleware
	playerMiddleware.Role = "player"
	r := e.Group("/v1/api")

	playerAuthRouter := r.Group("/auth")
	playerAuthRouter.POST("/login", c.PlayerAuthController.Login)
	playerAuthRouter.POST("/register", c.PlayerAuthController.Register)

	playerRouter := r.Group("/player")
	playerRouter.Use(playerMiddleware.VerifyRole)
	playerRouter.POST("/status", c.PlayersController.FindByToken)

	eventsRouter := r.Group("/events")
	eventsRouter.GET("", c.EventsController.Find)

	equipmentsRouter := r.Group("/equipments")
	equipmentsRouter.GET("", c.EquipmentsController.Find)


	// API ADMIN
	adminMiddleware := *c.JWTMiddleware
	adminMiddleware.Role = "admin"
	rAdmin := r.Group("-admin")
	
	adminAuthRouter := rAdmin.Group("/auth")
	adminAuthRouter.POST("/login", c.AdminAuthController.Login)

	adminRouter := rAdmin.Group("/admin")
	adminRouter.Use(adminMiddleware.VerifyRole)
	adminRouter.POST("/token", c.AdminController.FindByToken)
	
}
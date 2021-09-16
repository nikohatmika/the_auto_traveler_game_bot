package main

import (
	"log"
	"time"

	adminAuthUsecase "auto_traveler/bussiness/admin_auth"
	adminAuthController "auto_traveler/controller/admin_auth"

	playerUsecase "auto_traveler/bussiness/players"
	playerController "auto_traveler/controller/players"
	playerRepository "auto_traveler/driver/database/players"

	adminUsecase "auto_traveler/bussiness/admin"
	adminController "auto_traveler/controller/admin"
	adminRepository "auto_traveler/driver/database/admin"

	eventsUsecase "auto_traveler/bussiness/events"
	eventsController "auto_traveler/controller/events"
	eventsRepository "auto_traveler/driver/database/events"

	playerAuthUsecase "auto_traveler/bussiness/player_auth"
	playerAuthController "auto_traveler/controller/player_auth"

	"auto_traveler/app/middleware"
	_routes "auto_traveler/app/routers"
	_dbHelper "auto_traveler/driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configJWT := middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	configdb := _dbHelper.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}
	
	db := configdb.InitialDB()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	adminRepo := adminRepository.NewAdminRepository(db)
	adminUc := adminUsecase.NewAdminUsecase(timeoutContext, adminRepo)
	adminCtrl := adminController.NewAdminController(e, adminUc)

	adminAuthUc := adminAuthUsecase.NewAdminAuthUsecase(timeoutContext, adminRepo, &configJWT)
	adminAuthCtrl := adminAuthController.NewAdminAuthController(e, adminAuthUc)

	playerRepo := playerRepository.NewPlayersRepository(db)
	playerUc := playerUsecase.NewPlayerUsecase(timeoutContext, playerRepo)
	playerCtrl := playerController.NewPlayerController(e, playerUc)

	playerAuthUc := playerAuthUsecase.NewplayerAuthUsecase(timeoutContext, playerRepo, &configJWT)
	playerAuthCtrl := playerAuthController.NewplayerAuthController(e, playerAuthUc)

	eventsRepo := eventsRepository.NewEventsRepository(db)
	eventsUc := eventsUsecase.NewEventsUsecase(timeoutContext, eventsRepo)
	eventsCtrl := eventsController.NewEventsController(e, eventsUc)

	routesInit := _routes.ControllerList{
		JWTMiddleware:          &configJWT,
		AdminAuthController:    *adminAuthCtrl,
		AdminController:    	*adminCtrl,
		PlayerAuthController: 	*playerAuthCtrl,
		PlayersController:     	*playerCtrl,
		EventsController: 		*eventsCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
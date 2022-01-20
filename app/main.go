package main

import (
	"log"

	_config "sirclo/restapi/layered/config"
	_middlewares "sirclo/restapi/layered/delivery/middlewares"
	_routes "sirclo/restapi/layered/delivery/routes"
	_util "sirclo/restapi/layered/util"

	_authController "sirclo/restapi/layered/delivery/controllers/auth"
	_userController "sirclo/restapi/layered/delivery/controllers/user"

	_authRepo "sirclo/restapi/layered/repository/auth"
	_userRepo "sirclo/restapi/layered/repository/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	// configDB := config.AppConfig{
	// 	Database.Username : viper.GetString(`database.user`),
	// 	Database.Password: viper.GetString(`database.pass`),
	// 	Database.Host:     viper.GetString(`database.host`),
	// 	Database.Port:     viper.GetString(`database.port`),
	// 	Database.Name: viper.GetString(`database.name`),
	// }
	configDB := _config.GetConfig()
	db := _util.GetDBInstance(configDB)

	// fallback := config.AppConfig{}
	// config := config.GetConfig(&fallback)

	// db := util.GetDBInstance(config)
	defer db.Close()

	authRepo := _authRepo.New(db)
	// bookRepo := _bookRepo.New(db)
	// productRepo := _productRepo.New(db)
	userRepo := _userRepo.New(db)

	authController := _authController.New(authRepo)
	// bookController := _bookController.New(bookRepo)
	// productController := _productController.New(productRepo)
	userController := _userController.New(userRepo)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash(), _middlewares.CustomLogger())

	_routes.RegisterPath(e, authController, userController)

	e.Logger.Fatal(e.Start((":8080")))
}

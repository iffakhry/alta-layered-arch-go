package routes

import (
	_authController "sirclo/restapi/layered/delivery/controllers/auth"
	_userController "sirclo/restapi/layered/delivery/controllers/user"
	_middlewares "sirclo/restapi/layered/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	authController *_authController.AuthController,
	userController *_userController.UserController,
) {

	// Login
	e.POST("/login", authController.Login())

	// User
	e.GET("/users", userController.GetAll(), _middlewares.JWTMiddleware())
	e.GET("/users/:id", userController.Get(), _middlewares.JWTMiddleware())
	e.POST("/users", userController.Create())
	e.PUT("/users/:id", userController.Update(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userController.Delete(), _middlewares.JWTMiddleware())

}

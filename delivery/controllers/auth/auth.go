package auth

import (
	"net/http"
	_controllers "sirclo/restapi/layered/delivery/controllers"
	_entity "sirclo/restapi/layered/entity"
	_authRepo "sirclo/restapi/layered/repository/auth"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repository _authRepo.Auth
}

func New(auth _authRepo.Auth) *AuthController {
	return &AuthController{
		repository: auth,
	}
}

func (a AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		login := _entity.User{}

		if err := c.Bind(&login); err != nil {
			// code := http.StatusBadRequest
			return _controllers.NewErrorResponse(c, http.StatusBadRequest, "invalid data request")
			// return c.JSON(code, common.SimpleResponse(code, "binding failed", ""))
		}

		token, code := a.repository.Login(*login.Name, *login.Password)

		if code != http.StatusOK {
			return _controllers.NewErrorResponse(c, code, token)
			// return c.JSON(code, common.SimpleResponse(code, token, ""))
		}
		response := map[string]string{
			"token": token,
		}
		return _controllers.NewSuccessResponse(c, "update user success", response)
		// return c.JSON(code, common.SimpleResponse(code, "login success", token))
	}
}

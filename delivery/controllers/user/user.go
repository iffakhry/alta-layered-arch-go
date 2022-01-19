package user

import (
	"net/http"
	_controllers "sirclo/restapi/layered/delivery/controllers"
	_middlewares "sirclo/restapi/layered/delivery/middlewares"
	_entity "sirclo/restapi/layered/entity"
	_userRepo "sirclo/restapi/layered/repository/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repository _userRepo.User
}

func New(user _userRepo.User) *UserController {
	return &UserController{
		repository: user,
	}
}

func (uc UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// code := http.StatusOK

		if valid := _middlewares.ValidateToken(c); !valid {
			// code = http.StatusUnauthorized
			return _controllers.NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
			// return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		users, err := uc.repository.GetAll()

		var userResponseData []UserResponseFormat
		for _, value := range users {
			userResponseData = append(userResponseData, FromEntity(value))

		}
		if err != nil {
			// code = http.StatusInternalServerError
			return _controllers.NewErrorResponse(c, http.StatusInternalServerError, "failed to get all users")
			// return c.JSON(code, common.SimpleResponse(code, "get all users failed", nil))
		}

		if len(users) == 0 {
			return _controllers.NewErrorResponse(c, http.StatusInternalServerError, "empty users data")
			// return c.JSON(code, common.SimpleResponse(code, "users directory empty", nil))
		}

		return _controllers.NewSuccessResponse(c, "success get all users", userResponseData)
		// return c.JSON(code, common.SimpleResponse(code, "get all users success", users))
	}
}

func (uc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		// code := http.StatusOK

		if valid := _middlewares.ValidateToken(c); !valid {
			return _controllers.NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
			// code = http.StatusUnauthorized
			// return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return _controllers.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
			// code = http.StatusBadRequest
			// return c.JSON(code, common.SimpleResponse(code, "invalid user id", nil))
		}

		user, err := uc.repository.Get(id)

		if err != nil {
			return _controllers.NewErrorResponse(c, http.StatusInternalServerError, "failed to get user")
			// code = http.StatusInternalServerError
			// return c.JSON(code, common.SimpleResponse(code, "get user failed", nil))
		}

		// if user == (common.UserResponse{}) {
		// 	code = http.StatusBadRequest
		// 	return c.JSON(code, common.SimpleResponse(code, "user does not exist", nil))
		// }

		return _controllers.NewSuccessResponse(c, "success get user", FromEntity(user))
		// return c.JSON(code, common.SimpleResponse(code, "get user success", []common.UserResponse{user}))
	}
}

func (uc UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := _entity.User{}
		// userRequest := UserRequestFormat{}
		// code := http.StatusOK

		if err := c.Bind(&user); err != nil {
			return _controllers.NewErrorResponse(c, http.StatusBadRequest, "invalid data request")
			// code = http.StatusBadRequest
			// return c.JSON(code, common.SimpleResponse(code, "binding failed", nil))
		}

		id, err := uc.repository.Create(user)

		if err != nil {
			return _controllers.NewErrorResponse(c, http.StatusInternalServerError, "failed to create new user")
			// code = http.StatusInternalServerError
			// return c.JSON(code, common.SimpleResponse(code, "create user failed", nil))
		}

		user.Id = id
		return _controllers.NewSuccessResponse(c, "success create users", user)
		// return c.JSON(code, common.SimpleResponse(code, "create user success", []_entity.User{user}))
	}
}

func (uc UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// code := http.StatusOK

		if valid := _middlewares.ValidateToken(c); !valid {
			return _controllers.NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
			// code = http.StatusUnauthorized
			// return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return _controllers.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
			// code = http.StatusBadRequest
			// return c.JSON(code, common.SimpleResponse(code, "invalid user id", nil))
		}

		// user := _entity.User{}
		userRequestData := UserRequestFormat{}

		if err := c.Bind(&userRequestData); err != nil {
			return _controllers.NewErrorResponse(c, http.StatusBadRequest, "invalid data request")
			// code = http.StatusBadRequest
			// return c.JSON(code, common.SimpleResponse(code, "binding failed", nil))
		}

		var userData = userRequestData.ToEntity()
		userData.Id = id

		if code, err := uc.repository.Update(*userData); err != nil {
			return _controllers.NewErrorResponse(c, code, err.Error())
			// return c.JSON(code, common.SimpleResponse(code, err.Error(), nil))
		}

		return _controllers.NewSuccessResponse(c, "update user success", userData)
		// return c.JSON(code, common.SimpleResponse(code, "update user success", []_entity.User{user}))
	}
}

func (uc UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// code := http.StatusOK

		if valid := _middlewares.ValidateToken(c); !valid {
			return _controllers.NewErrorResponse(c, http.StatusUnauthorized, "unauthorized")
			// code = http.StatusUnauthorized
			// return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return _controllers.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
			// code = http.StatusBadRequest
			// return c.JSON(code, common.SimpleResponse(code, "invalid user id", nil))
		}

		if code, err := uc.repository.Delete(id); err != nil {
			return _controllers.NewErrorResponse(c, code, err.Error())
			// return c.JSON(code, common.SimpleResponse(code, err.Error(), nil))
		}

		return _controllers.NewSuccessResponse(c, "delete user success", nil)
		// return c.JSON(code, common.SimpleResponse(code, "delete user success", nil))
	}
}

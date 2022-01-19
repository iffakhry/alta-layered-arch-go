package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const secret_jwt = "Xcx*3-bPr9w&USU7"

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(secret_jwt),
	})
}

func CreateToken(id int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret_jwt))
}

func ValidateToken(e echo.Context) bool {
	login := e.Get("user").(*jwt.Token)

	return login.Valid
}

func ExtractId(e echo.Context) (int, error) {
	login := e.Get("user").(*jwt.Token)

	if login.Valid {
		claims := login.Claims.(jwt.MapClaims)
		id := int(claims["id"].(float64))
		return id, nil
	}

	return 0, fmt.Errorf("unauthorized")
}

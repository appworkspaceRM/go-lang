package middleware

import (
	dtoressult "landtick_backend/dto/ressult"
	"landtick_backend/pkg/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dtoressult.ErrorResult{
				Status: "Error",
				Message: "Unauthorized",
			})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwt.DecodeToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dtoressult.ErrorResult{
				Status: "Error",
				Message: "Unauthorized",
			})
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}
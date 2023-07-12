package handlerauth

import (
	dtoauth "landtick_backend/dto/auth"
	dtoressult "landtick_backend/dto/ressult"
	repositoryauth "landtick_backend/repositories/auth"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositoryauth.AuthRepository
}

type JsonAuth struct {
	DataAuth interface{} `json:"user"`
}

func HandlerAuth(AuthRepository repositoryauth.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth)Register(c echo.Context) error  {
	request := new(dtoauth.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dtoressult.ErrorResult{
			Status: "Error request",
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoressult.ErrorResult{
			Status: "Error request",
			Message: err.Error(),
		})
	}
	return nil
}
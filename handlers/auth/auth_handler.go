package handlerauth

import (
	dtoauth "landtick_backend/dto/auth"
	dtoressult "landtick_backend/dto/ressult"
	modeluser "landtick_backend/models/user"
	pkgbcrypt "landtick_backend/pkg/bcrypt"
	repositoryauth "landtick_backend/repositories/auth"
	"net/http"
	"time"

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

	password, err := pkgbcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoressult.ErrorResult{
			Status: "Error",
			Message: err.Error(),
		})
	}

	user := modeluser.User{
		FullName: request.FullName,
		UserName: request.UserName,
		Email: request.Email,
		Password: password,
		Role: "customer",
		PhoneNumber: request.PhoneNumber,
		Gender: request.Gender,
		Address: request.Address,
		CreatedAd: time.Now(),
		UpdateAd: time.Now(),
	}

	register, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoressult.ErrorResult{
			Status: "Error",
			Message: err.Error(),
		})
	}	
	return c.JSON(http.StatusOK, dtoressult.SuccessResult{
		Stattus: "success",
		Data: JsonAuth{
			DataAuth: register,
		},
	})
}
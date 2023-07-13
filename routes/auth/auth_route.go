package routeauth

import (
	handlerauth "landtick_backend/handlers/auth"
	"landtick_backend/pkg/mysql"
	repositoryauth "landtick_backend/repositories/auth"

	"github.com/labstack/echo/v4"
)

func AuthRoute(e *echo.Group) {
	authRepository := repositoryauth.RepositoryAuth(mysql.DB)
	h := handlerauth.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
}
package routes

import (
	routeauth "landtick_backend/routes/auth"

	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Group) {
	routeauth.AuthRoute(e)
}
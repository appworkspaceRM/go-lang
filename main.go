package main

import (
	"landtick_backend/database"
	"landtick_backend/pkg/mysql"
	"landtick_backend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start("localhost:5000"))
}
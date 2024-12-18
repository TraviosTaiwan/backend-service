package main

import (
	"github.com/iwachan14736/travios-backend-service/pkg/containers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	containers.Serve(e)
}

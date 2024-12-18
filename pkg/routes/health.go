package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthRoutes(group *echo.Group) {
	group.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})
}

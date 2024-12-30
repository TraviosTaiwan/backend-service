package routes

import (
	"github.com/iwachan14736/travios-backend-service/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(group *echo.Group) {
	group.POST("/login", controllers.Login)
	group.POST("/register", controllers.Register)
	group.GET("/anonymous-token", controllers.GetAnonymousToken)
}

package routes

import (
	"github.com/iwachan14736/travios-backend-service/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func ItemRoutes(group *echo.Group) {
	group.POST("/item", controllers.CreateItem)
	group.GET("/item", controllers.GetItems)
	group.PUT("/item/:itemID", controllers.UpdateItem)
	group.DELETE("/item/:itemID", controllers.DeleteItem)
}

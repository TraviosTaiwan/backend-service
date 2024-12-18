package routes

import (
	"github.com/iwachan14736/travios-backend-service/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func CategoryRoutes(group *echo.Group) {
	group.POST("/category", controllers.CreateCategory)
	group.GET("/category", controllers.GetCategories)
	group.PUT("/category/:categoryID", controllers.UpdateCategory)
	group.DELETE("/category/:categoryID", controllers.DeleteCategory)
}

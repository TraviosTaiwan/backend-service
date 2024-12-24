package routes

import (
	"github.com/iwachan14736/travios-backend-service/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func CustomerRoutes(group *echo.Group) {
	group.POST("/customer", controllers.CreateCustomer)
	group.GET("/customer", controllers.GetCustomers)
	group.PUT("/customer/:customerID", controllers.UpdateCustomer)
	group.DELETE("/customer/:customerID", controllers.DeleteCustomer)
}

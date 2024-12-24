package routes

import (
	"github.com/iwachan14736/travios-backend-service/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func SaleOrderRoutes(group *echo.Group) {
	group.POST("/sale_order", controllers.CreateSaleOrder)
	group.GET("/sale_order", controllers.GetSaleOrders)
	group.PUT("/sale_order/:sale_orderID", controllers.UpdateSaleOrder)
	group.DELETE("/sale_order/:sale_orderID", controllers.DeleteSaleOrder)
}

package routes

import (
	"github.com/iwachan14736/travios-backend-service/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func SaleOrderItemRoutes(group *echo.Group) {
	group.POST("/sale_order_item", controllers.CreateSaleOrderItem)
	group.GET("/sale_order_item", controllers.GetSaleOrderItems)
	group.PUT("/sale_order_item/:sale_order_itemID", controllers.UpdateSaleOrderItem)
	group.DELETE("/sale_order_item/:sale_order_itemID", controllers.DeleteSaleOrderItem)
}

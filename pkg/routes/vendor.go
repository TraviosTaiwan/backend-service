package routes

import (
	"github.com/iwachan14736/travios-backend-service/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func VendorRoutes(group *echo.Group) {
	group.POST("/vendor", controllers.CreateVendor)
	group.GET("/vendor", controllers.GetVendors)
	group.PUT("/vendor/:vendorID", controllers.UpdateVendor)
	group.DELETE("/vendor/:vendorID", controllers.DeleteVendor)
}

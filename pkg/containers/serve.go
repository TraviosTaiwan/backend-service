package containers

import (
	"fmt"
	"log"

	"github.com/iwachan14736/travios-backend-service/pkg/config"
	"github.com/iwachan14736/travios-backend-service/pkg/controllers"
	"github.com/iwachan14736/travios-backend-service/pkg/database"
	"github.com/iwachan14736/travios-backend-service/pkg/repositories"
	"github.com/iwachan14736/travios-backend-service/pkg/routes"
	"github.com/iwachan14736/travios-backend-service/pkg/services"
	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	config.SetConfig()
	db := database.GetDB()

	// Repositories
	userRepository := repositories.UserDBInstance(db)
	categoryRepository := repositories.CategoryDBInstance(db)
	vendorRepository := repositories.VendorDBInstance(db)
	itemRepository := repositories.ItemDBInstance(db)

	// Services
	authService := services.AuthServiceInstance(userRepository)
	categoryService := services.CategoryServiceInstance(categoryRepository)
	vendorService := services.VendorServiceInstance(vendorRepository)
	itemService := services.ItemServiceInstance(itemRepository)

	// Controllers
	controllers.SetAuthService(authService)
	controllers.SetCategoryService(categoryService)
	controllers.SetVendorService(vendorService)
	controllers.SetItemService(itemService)
	routes.InitRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}

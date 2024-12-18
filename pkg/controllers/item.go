package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
	"github.com/labstack/echo/v4"
)

var ItemService domain.IItemService

func SetItemService(vService domain.IItemService) {
	ItemService = vService
}

// @Summary Create a new item
// @Description Create a new item with the provided details
// @Tags items
// @Accept json
// @Produce json
// @Security Bearer
// @Param item body types.Item true "Item"
// @Success 201 {object} types.Item
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /item [post]
func CreateItem(e echo.Context) error {
	reqItem := &types.Item{}
	if err := e.Bind(reqItem); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqItem.ValidateItem(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	log.Println(reqItem)
	item := &models.Item{
		Name:         reqItem.Name,
		CategoryID:   reqItem.CategoryID,
		VendorID:     reqItem.VendorID,
		Description:  reqItem.Description,
		ImageUrl:     reqItem.ImageUrl,
		BuyingPrice:  reqItem.BuyingPrice,
		SellingPrice: reqItem.SellingPrice,
		Quantity:     reqItem.Quantity,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := ItemService.CreateItem(item); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, item)
}

// @Summary Get items
// @Description Get a list of items or a specific item by ID
// @Tags items
// @Produce json
// @Security Bearer
// @Param itemID query int false "Item ID"
// @Success 200 {array} types.Item
// @Failure 401 {string} string "Unauthorized"
// @Router /item [get]
func GetItems(e echo.Context) error {
	itemID, _ := strconv.Atoi(e.QueryParam("itemID"))
	items, err := ItemService.GetItems(uint(itemID))
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}
	return e.JSON(http.StatusOK, items)
}

// @Summary Update an item
// @Description Update an item with the provided details
// @Tags items
// @Accept json
// @Produce json
// @Security Bearer
// @Param itemID path int true "Item ID"
// @Param item body types.Item true "Item"
// @Success 200 {object} types.Item
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /item/{itemID} [put]
func UpdateItem(e echo.Context) error {
	itemID, _ := strconv.Atoi(e.Param("itemID"))
	reqItem := &types.Item{}
	if err := e.Bind(reqItem); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqItem.ValidateItem(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	item := &models.Item{
		ID:           uint(itemID),
		Name:         reqItem.Name,
		CategoryID:   reqItem.CategoryID,
		VendorID:     reqItem.VendorID,
		BuyingPrice:  reqItem.BuyingPrice,
		SellingPrice: reqItem.SellingPrice,
		Description:  reqItem.Description,
		ImageUrl:     reqItem.ImageUrl,
		Quantity:     reqItem.Quantity,
	}
	if err := ItemService.UpdateItem(item); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, item)
}

// @Summary Delete an item
// @Description Delete an item by ID
// @Tags items
// @Produce json
// @Security Bearer
// @Param itemID path int true "Item ID"
// @Success 200 {string} string "Item deleted successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Item not found"
// @Router /item/{itemID} [delete]
func DeleteItem(e echo.Context) error {
	itemID, _ := strconv.Atoi(e.Param("itemID"))
	if err := ItemService.DeleteItem(uint(itemID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Item deleted successfully")
}

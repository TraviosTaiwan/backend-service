package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
	"github.com/labstack/echo/v4"
)

var SaleOrderItemService domain.ISaleOrderItemService

func SetSaleOrderItemService(vService domain.ISaleOrderItemService) {
	SaleOrderItemService = vService
}

// @Summary Create a new sale order item
// @Description Create a new sale order item with the provided details
// @Tags sale_order_items
// @Accept json
// @Produce json
// @Security Bearer
// @Param sale_order_item body types.SaleOrderItem true "Sale Order Item"
// @Success 201 {object} types.SaleOrderItem
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /sale_order_item [post]
func CreateSaleOrderItem(e echo.Context) error {
	reqSaleOrderItem := &types.SaleOrderItem{}
	if err := e.Bind(reqSaleOrderItem); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqSaleOrderItem.ValidateSaleOrderItem(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	log.Println(reqSaleOrderItem)
	saleOrderItem := &models.SaleOrderItem{
		SaleOrderID: reqSaleOrderItem.SaleOrderID,
		ItemID:      reqSaleOrderItem.ItemID,
		Quantity:    reqSaleOrderItem.Quantity,
	}
	if err := SaleOrderItemService.CreateSaleOrderItem(saleOrderItem); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, saleOrderItem)
}

// @Summary Get sale order items
// @Description Get a list of sale order items or a specific sale order item by ID
// @Tags sale_order_items
// @Produce json
// @Security Bearer
// @Param sale_order_itemID query int false "Sale Order Item ID"
// @Success 200 {array} types.SaleOrderItem
// @Failure 401 {string} string "Unauthorized"
// @Router /sale_order_item [get]
func GetSaleOrderItems(e echo.Context) error {
	saleOrderItemID, _ := strconv.Atoi(e.QueryParam("sale_order_itemID"))
	saleOrderItems, err := SaleOrderItemService.GetSaleOrderItems(uint(saleOrderItemID))
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}
	return e.JSON(http.StatusOK, saleOrderItems)
}

// @Summary Get sale order items by sale order ID
// @Description Get a list of sale order items by sale order ID
// @Tags sale_order_items
// @Produce json
// @Security Bearer
// @Param sale_order_id query int true "Sale Order ID"
// @Success 200 {array} types.SaleOrderItem
// @Failure 401 {string} string "Unauthorized"
// @Router /sale_order_item/sale_order [get]
func GetSaleOrderItemsBySaleOrderID(e echo.Context) error {
	saleOrderID, _ := strconv.Atoi(e.QueryParam("sale_order_id"))
	saleOrderItems, err := SaleOrderItemService.GetSaleOrderItemsBySaleOrderID(uint(saleOrderID))
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}
	return e.JSON(http.StatusOK, saleOrderItems)
}

// @Summary Update an sale order item
// @Description Update an sale order item with the provided details
// @Tags sale_order_items
// @Accept json
// @Produce json
// @Security Bearer
// @Param sale_order_itemID path int true "Sale Order Item ID"
// @Param sale_order_item body types.SaleOrderItem true "Sale Order Item"
// @Success 200 {object} types.SaleOrderItem
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /sale_order_item/{sale_order_itemID} [put]
func UpdateSaleOrderItem(e echo.Context) error {
	saleOrderItemID, _ := strconv.Atoi(e.Param("sale_order_itemID"))
	reqSaleOrderItem := &types.SaleOrderItem{}
	if err := e.Bind(reqSaleOrderItem); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqSaleOrderItem.ValidateSaleOrderItem(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	saleOrderItem := &models.SaleOrderItem{
		ID:          uint(saleOrderItemID),
		SaleOrderID: reqSaleOrderItem.SaleOrderID,
		ItemID:      reqSaleOrderItem.ItemID,
		Quantity:    reqSaleOrderItem.Quantity,
	}
	if err := SaleOrderItemService.UpdateSaleOrderItem(saleOrderItem); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, saleOrderItem)
}

// @Summary Delete an sale order item
// @Description Delete an sale order item by ID
// @Tags sale_order_items
// @Produce json
// @Security Bearer
// @Param sale_order_itemID path int true "Sale Order Item ID"
// @Success 200 {string} string "Sale Order Item deleted successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Sale Order Item not found"
// @Router /sale_order_item/{sale_order_itemID} [delete]
func DeleteSaleOrderItem(e echo.Context) error {
	saleOrderItemID, _ := strconv.Atoi(e.Param("sale_order_itemID"))
	if err := SaleOrderItemService.DeleteSaleOrderItem(uint(saleOrderItemID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Sale Order Item deleted successfully")
}

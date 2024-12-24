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

var SaleOrderService domain.ISaleOrderService

func SetSaleOrderService(vService domain.ISaleOrderService) {
	SaleOrderService = vService
}

// @Summary Create a new sale order
// @Description Create a new sale order with the provided details
// @Tags sale_orders
// @Accept json
// @Produce json
// @Security Bearer
// @Param sale_order body types.SaleOrder true "Sale Order"
// @Success 201 {object} types.SaleOrder
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /sale_order [post]
func CreateSaleOrder(e echo.Context) error {
	reqSaleOrder := &types.SaleOrder{}
	if err := e.Bind(reqSaleOrder); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqSaleOrder.ValidateSaleOrder(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	log.Println(reqSaleOrder)
	saleOrder := &models.SaleOrder{
		CustomerID: reqSaleOrder.CustomerID,
	}
	if err := SaleOrderService.CreateSaleOrder(saleOrder); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, saleOrder)
}

// @Summary Get sale orders
// @Description Get a list of sale orders or a specific sale order by ID
// @Tags sale_orders
// @Produce json
// @Security Bearer
// @Param sale_orderID query int false "Sale Order ID"
// @Success 200 {array} types.SaleOrder
// @Failure 401 {string} string "Unauthorized"
// @Router /sale_order [get]
func GetSaleOrders(e echo.Context) error {
	saleOrderID, _ := strconv.Atoi(e.QueryParam("sale_orderID"))
	saleOrders, err := SaleOrderService.GetSaleOrders(uint(saleOrderID))
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}
	return e.JSON(http.StatusOK, saleOrders)
}

// @Summary Update an sale order
// @Description Update an sale order with the provided details
// @Tags sale_orders
// @Accept json
// @Produce json
// @Security Bearer
// @Param sale_orderID path int true "Sale Order ID"
// @Param sale_order body types.SaleOrder true "Sale Order"
// @Success 200 {object} types.SaleOrder
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /sale_order/{sale_orderID} [put]
func UpdateSaleOrder(e echo.Context) error {
	saleOrderID, _ := strconv.Atoi(e.Param("sale_orderID"))
	reqSaleOrder := &types.SaleOrder{}
	if err := e.Bind(reqSaleOrder); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqSaleOrder.ValidateSaleOrder(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	saleOrder := &models.SaleOrder{
		ID:         uint(saleOrderID),
		CustomerID: reqSaleOrder.CustomerID,
	}
	if err := SaleOrderService.UpdateSaleOrder(saleOrder); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, saleOrder)
}

// @Summary Delete an sale order
// @Description Delete an sale order by ID
// @Tags sale_orders
// @Produce json
// @Security Bearer
// @Param sale_orderID path int true "Sale Order ID"
// @Success 200 {string} string "Sale Order deleted successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Sale Order not found"
// @Router /sale_order/{sale_orderID} [delete]
func DeleteSaleOrder(e echo.Context) error {
	saleOrderID, _ := strconv.Atoi(e.Param("sale_orderID"))
	if err := SaleOrderService.DeleteSaleOrder(uint(saleOrderID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Sale Order deleted successfully")
}

package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/iwachan14736/travios-backend-service/pkg/database"
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
		Remark:     reqSaleOrder.Remark,
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
		Remark:     reqSaleOrder.Remark,
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

// @Summary Create a new sale order from frontend
// @Description Create a new sale order with customer info and cart items
// @Tags sale_orders
// @Accept json
// @Produce json
// @Param sale_order body types.SaleOrderRequest true "Sale Order Request"
// @Param Authorization header string false "Anonymous token (optional)"
// @Success 201 {object} types.SaleOrder
// @Failure 400 {string} string "Invalid input"
// @Router /sale_order/frontend [post]
func CreateSaleOrderFromFrontend(e echo.Context) error {
	reqSaleOrder := &types.SaleOrderRequest{}
	if err := e.Bind(reqSaleOrder); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}

	if err := reqSaleOrder.ValidateSaleOrderRequest(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	// Start transaction
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find or create customer with platform "web"
	customer := &models.Customer{
		Name:     reqSaleOrder.CustomerName,
		Phone:    reqSaleOrder.CustomerPhone,
		Address:  reqSaleOrder.CustomerAddress,
		Platform: "web", // Auto set platform to "web" for frontend requests
	}

	// Try to find existing customer by phone number
	var existingCustomer models.Customer
	if err := tx.Where("phone = ?", customer.Phone).First(&existingCustomer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new customer if not found
			if err := tx.Create(customer).Error; err != nil {
				tx.Rollback()
				return e.JSON(http.StatusInternalServerError, "Failed to create customer")
			}
		} else {
			tx.Rollback()
			return e.JSON(http.StatusInternalServerError, "Database error")
		}
	} else {
		// Update existing customer's platform if it's not already "web"
		if existingCustomer.Platform != "web" {
			existingCustomer.Platform = "web"
			if err := tx.Save(&existingCustomer).Error; err != nil {
				tx.Rollback()
				return e.JSON(http.StatusInternalServerError, "Failed to update customer platform")
			}
		}
		customer = &existingCustomer
	}

	// Create sale order
	saleOrder := &models.SaleOrder{
		CustomerID: customer.ID,
		Remark:     reqSaleOrder.Remark,
	}

	if err := tx.Create(saleOrder).Error; err != nil {
		tx.Rollback()
		return e.JSON(http.StatusInternalServerError, "Failed to create sale order")
	}

	// Create sale order items
	for _, item := range reqSaleOrder.CartItems {
		saleOrderItem := &models.SaleOrderItem{
			SaleOrderID: saleOrder.ID,
			ItemID:      item.ItemID,
			Quantity:    item.Quantity,
		}

		if err := tx.Create(saleOrderItem).Error; err != nil {
			tx.Rollback()
			return e.JSON(http.StatusInternalServerError, "Failed to create sale order items")
		}

		// Update item quantity
		if err := tx.Model(&models.Item{}).Where("id = ?", item.ItemID).
			UpdateColumn("quantity", gorm.Expr("quantity - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return e.JSON(http.StatusInternalServerError, "Failed to update item quantity")
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return e.JSON(http.StatusInternalServerError, "Failed to commit transaction")
	}

	return e.JSON(http.StatusCreated, saleOrder)
}

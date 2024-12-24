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

var CustomerService domain.ICustomerService

func SetCustomerService(vService domain.ICustomerService) {
	CustomerService = vService
}

// @Summary Create a new customer
// @Description Create a new customer with the provided details
// @Tags customers
// @Accept json
// @Produce json
// @Security Bearer
// @Param customer body types.Customer true "Customer"
// @Success 201 {object} types.Customer
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /customer [post]
func CreateCustomer(e echo.Context) error {
	reqCustomer := &types.Customer{}
	if err := e.Bind(reqCustomer); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqCustomer.ValidateCustomer(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	log.Println(reqCustomer)
	customer := &models.Customer{
		Phone:    reqCustomer.Phone,
		Platform: reqCustomer.Platform,
		Username: reqCustomer.Username,
		Name:     reqCustomer.Name,
		Address:  reqCustomer.Address,
	}
	if err := CustomerService.CreateCustomer(customer); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, customer)
}

// @Summary Get customers
// @Description Get a list of customers or a specific customer by ID
// @Tags customers
// @Produce json
// @Security Bearer
// @Param customerID query int false "Customer ID"
// @Success 200 {array} types.Customer
// @Failure 401 {string} string "Unauthorized"
// @Router /customer [get]
func GetCustomers(e echo.Context) error {
	customerID, _ := strconv.Atoi(e.QueryParam("customerID"))
	customers, err := CustomerService.GetCustomers(uint(customerID))
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}
	return e.JSON(http.StatusOK, customers)
}

// @Summary Update an customer
// @Description Update an customer with the provided details
// @Tags customers
// @Accept json
// @Produce json
// @Security Bearer
// @Param customerID path int true "Customer ID"
// @Param customer body types.Customer true "Customer"
// @Success 200 {object} types.Customer
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /customer/{customerID} [put]
func UpdateCustomer(e echo.Context) error {
	customerID, _ := strconv.Atoi(e.Param("customerID"))
	reqCustomer := &types.Customer{}
	if err := e.Bind(reqCustomer); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqCustomer.ValidateCustomer(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	customer := &models.Customer{
		ID:       uint(customerID),
		Phone:    reqCustomer.Phone,
		Platform: reqCustomer.Platform,
		Username: reqCustomer.Username,
		Name:     reqCustomer.Name,
		Address:  reqCustomer.Address,
	}
	if err := CustomerService.UpdateCustomer(customer); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, customer)
}

// @Summary Delete an customer
// @Description Delete an customer by ID
// @Tags customers
// @Produce json
// @Security Bearer
// @Param customerID path int true "Customer ID"
// @Success 200 {string} string "Customer deleted successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Customer not found"
// @Router /customer/{customerID} [delete]
func DeleteCustomer(e echo.Context) error {
	customerID, _ := strconv.Atoi(e.Param("customerID"))
	if err := CustomerService.DeleteCustomer(uint(customerID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Customer deleted successfully")
}

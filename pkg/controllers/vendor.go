package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
	"github.com/labstack/echo/v4"
)

var VendorService domain.IVendorService

func SetVendorService(vService domain.IVendorService) {
	VendorService = vService
}

// @Summary Create a new vendor
// @Description Create a new vendor with the provided details
// @Tags vendors
// @Accept json
// @Produce json
// @Security Bearer
// @Param vendor body types.Vendor true "Vendor"
// @Success 201 {object} types.Vendor
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /vendor [post]
func CreateVendor(e echo.Context) error {
	reqVendor := &types.Vendor{}
	if err := e.Bind(reqVendor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqVendor.ValidateVendor(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	vendor := &models.Vendor{
		Name:      reqVendor.Name,
		Phone:     reqVendor.Phone,
		Address:   reqVendor.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := VendorService.CreateVendor(vendor); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, vendor)
}

// @Summary Get vendors
// @Description Get a list of vendors or a specific vendor by ID
// @Tags vendors
// @Produce json
// @Security Bearer
// @Param vendorID query int false "Vendor ID"
// @Success 200 {array} types.Vendor
// @Failure 401 {string} string "Unauthorized"
// @Router /vendor [get]
func GetVendors(e echo.Context) error {
	vendorID, _ := strconv.Atoi(e.QueryParam("vendorID"))
	vendors, err := VendorService.GetVendors(uint(vendorID))
	if err != nil {
		return e.JSON(http.StatusNotFound, err.Error())
	}
	return e.JSON(http.StatusOK, vendors)
}

// @Summary Update a vendor
// @Description Update a vendor with the provided details
// @Tags vendors
// @Accept json
// @Produce json
// @Security Bearer
// @Param vendorID path int true "Vendor ID"
// @Param vendor body types.Vendor true "Vendor"
// @Success 200 {object} types.Vendor
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Router /vendor/{vendorID} [put]
func UpdateVendor(e echo.Context) error {
	vendorID, _ := strconv.Atoi(e.Param("vendorID"))
	reqVendor := &types.Vendor{}
	if err := e.Bind(reqVendor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqVendor.ValidateVendor(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	vendor := &models.Vendor{
		ID:      uint(vendorID),
		Name:    reqVendor.Name,
		Phone:   reqVendor.Phone,
		Address: reqVendor.Address,
	}
	if err := VendorService.UpdateVendor(vendor); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, vendor)
}

// @Summary Delete a vendor
// @Description Delete a vendor by ID
// @Tags vendors
// @Produce json
// @Security Bearer
// @Param vendorID path int true "Vendor ID"
// @Success 200 {string} string "Vendor deleted successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Vendor not found"
// @Router /vendor/{vendorID} [delete]
func DeleteVendor(e echo.Context) error {
	vendorID, _ := strconv.Atoi(e.Param("vendorID"))
	if err := VendorService.DeleteVendor(uint(vendorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Vendor deleted successfully")
}

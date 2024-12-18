package domain

import (
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type IVendorRepository interface {
	GetVendors(vendorID uint) []models.Vendor
	CreateVendor(vendor *models.Vendor) error
	UpdateVendor(vendor *models.Vendor) error
	DeleteVendor(vendorID uint) error
}

type IVendorService interface {
	GetVendors(vendorID uint) ([]types.Vendor, error)
	CreateVendor(vendor *models.Vendor) error
	UpdateVendor(vendor *models.Vendor) error
	DeleteVendor(vendorID uint) error
}

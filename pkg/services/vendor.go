package services

import (
	"errors"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type VendorService struct {
	repo domain.IVendorRepository
}

func VendorServiceInstance(vendorRepo domain.IVendorRepository) domain.IVendorService {
	return &VendorService{
		repo: vendorRepo,
	}
}

func (service *VendorService) GetVendors(vendorID uint) ([]types.Vendor, error) {
	var allVendors []types.Vendor
	vendors := service.repo.GetVendors(vendorID)
	if len(vendors) == 0 {
		return nil, errors.New("Vendor not found")
	}
	for _, val := range vendors {
		allVendors = append(allVendors, types.Vendor{
			ID:      val.ID,
			Name:    val.Name,
			Phone:   val.Phone,
			Address: val.Address,
		})
	}
	return allVendors, nil
}

func (service *VendorService) CreateVendor(vendorRequest *models.Vendor) error {
	return service.repo.CreateVendor(vendorRequest)
}

func (service *VendorService) UpdateVendor(vendorRequest *models.Vendor) error {
	return service.repo.UpdateVendor(vendorRequest)
}

func (service *VendorService) DeleteVendor(vendorID uint) error {
	return service.repo.DeleteVendor(vendorID)
}

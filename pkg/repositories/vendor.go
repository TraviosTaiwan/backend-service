package repositories

import (
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"gorm.io/gorm"
)

type vendorRepo struct {
	db *gorm.DB
}

func VendorDBInstance(d *gorm.DB) domain.IVendorRepository {
	return &vendorRepo{
		db: d,
	}
}

func (repo *vendorRepo) CreateVendor(vendor *models.Vendor) error {
	if err := repo.db.Create(vendor).Error; err != nil {
		return err
	}
	return nil
}

func (repo *vendorRepo) GetVendors(vendorID uint) []models.Vendor {
	var vendors []models.Vendor
	var err error

	if vendorID != 0 {
		err = repo.db.Where("id = ?", vendorID).Find(&vendors).Error
	} else {
		err = repo.db.Find(&vendors).Error
	}

	if err != nil {
		return []models.Vendor{}
	}
	return vendors
}

func (repo *vendorRepo) UpdateVendor(vendor *models.Vendor) error {
	if err := repo.db.Save(vendor).Error; err != nil {
		return err
	}
	return nil
}

func (repo *vendorRepo) DeleteVendor(vendorID uint) error {
	var vendor models.Vendor
	if err := repo.db.
		Where("id = ?", vendorID).Delete(&vendor).
		Error; err != nil {
		return err
	}
	return nil
}

package repositories

import (
	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func CustomerDBInstance(d *gorm.DB) domain.ICustomerRepository {
	return &customerRepo{
		db: d,
	}
}

func (repo *customerRepo) CreateCustomer(customer *models.Customer) error {
	if err := repo.db.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *customerRepo) GetCustomers(customerID uint) []models.Customer {
	var customers []models.Customer
	var err error

	if customerID != 0 {
		err = repo.db.Where("id = ?", customerID).Find(&customers).Error
	} else {
		err = repo.db.Find(&customers).Error
	}

	if err != nil {
		return []models.Customer{}
	}
	return customers
}

func (repo *customerRepo) UpdateCustomer(customer *models.Customer) error {
	if err := repo.db.Save(customer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *customerRepo) DeleteCustomer(customerID uint) error {
	var customer models.Customer
	if err := repo.db.
		Where("id = ?", customerID).Delete(&customer).
		Error; err != nil {
		return err
	}
	return nil
}

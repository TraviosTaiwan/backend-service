package domain

import (
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type ICustomerRepository interface {
	GetCustomers(customerID uint) []models.Customer
	CreateCustomer(customer *models.Customer) error
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(customerID uint) error
}

type ICustomerService interface {
	GetCustomers(customerID uint) ([]types.Customer, error)
	CreateCustomer(customer *models.Customer) error
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(customerID uint) error
}

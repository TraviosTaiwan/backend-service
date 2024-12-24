package services

import (
	"errors"

	"github.com/iwachan14736/travios-backend-service/pkg/domain"
	"github.com/iwachan14736/travios-backend-service/pkg/models"
	"github.com/iwachan14736/travios-backend-service/pkg/types"
)

type CustomerService struct {
	repo domain.ICustomerRepository
}

func CustomerServiceInstance(customerRepo domain.ICustomerRepository) domain.ICustomerService {
	return &CustomerService{
		repo: customerRepo,
	}
}

func (service *CustomerService) GetCustomers(customerID uint) ([]types.Customer, error) {
	var allCustomers []types.Customer
	customers := service.repo.GetCustomers(customerID)
	if len(customers) == 0 {
		return nil, errors.New("Customer not found")
	}
	for _, val := range customers {
		allCustomers = append(allCustomers, types.Customer{
			ID:       val.ID,
			Name:     val.Name,
			Phone:    val.Phone,
			Address:  val.Address,
			Platform: val.Platform,
			Username: val.Username,
		})
	}
	return allCustomers, nil
}

func (service *CustomerService) CreateCustomer(customerRequest *models.Customer) error {
	return service.repo.CreateCustomer(customerRequest)
}

func (service *CustomerService) UpdateCustomer(customerRequest *models.Customer) error {
	return service.repo.UpdateCustomer(customerRequest)
}

func (service *CustomerService) DeleteCustomer(customerID uint) error {
	return service.repo.DeleteCustomer(customerID)
}

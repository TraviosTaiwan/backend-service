package types

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Customer struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Platform string `json:"platform"`
	Username string `json:"username"`
}

func (customer Customer) ValidateCustomer() error {
	return validation.ValidateStruct(&customer,
		validation.Field(&customer.Phone,
			validation.Required.Error("Phone cannot be empty"),
		),
	)
}

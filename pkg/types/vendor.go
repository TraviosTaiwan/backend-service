package types

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Vendor struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (vendor Vendor) ValidateVendor() error {
	return validation.ValidateStruct(&vendor,
		validation.Field(&vendor.Name,
			validation.Required.Error("Name cannot be empty"),
			validation.Length(1, 50),
		),
	)
}

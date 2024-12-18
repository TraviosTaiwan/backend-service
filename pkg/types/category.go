package types

import (
	validation "github.com/Go-ozzo/ozzo-validation"
)

type Category struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (category Category) ValidateCategory() error {
	return validation.ValidateStruct(&category,
		validation.Field(&category.Name,
			validation.Required.Error("Name cannot be empty"),
			validation.Length(1, 50),
		),
	)
}

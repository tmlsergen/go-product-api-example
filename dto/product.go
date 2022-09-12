package dto

import (
	"app/models"
	"github.com/go-playground/validator/v10"
)

type ProductRequestDto struct {
	Name        string  `json:"name,omitempty" validate:"required"`
	Price       float32 `json:"price,omitempty" validate:"required,number"`
	Description string  `json:"description,omitempty" validate:"required"`
	Currency    string  `json:"currency,omitempty" validate:"required"`
	SellerId    string  `json:"seller_id,omitempty" validate:"required"`
	InStock     bool    `json:"in_stock" validate:"required"`
}

type ProductResponseDto struct {
	Name            string                  `json:"name,omitempty"`
	Price           float32                 `json:"price,omitempty"`
	Description     string                  `json:"description,omitempty"`
	Currency        string                  `json:"currency,omitempty"`
	SellerId        string                  `json:"seller_id,omitempty"`
	InStock         bool                    `json:"in_stock"`
	DeliveryOptions []models.DeliveryOption `json:"delivery_options"`
}

func ValidateProductStruct(product ProductRequestDto) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(product)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

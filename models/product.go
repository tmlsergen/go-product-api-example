package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name            string           `json:"name,omitempty" gorm:"type:varchar(100)"`
	Price           float32          `json:"price,omitempty" gorm:"index;sort:desc;"`
	Description     string           `json:"description,omitempty" gorm:"type:varchar(100)"`
	Currency        string           `json:"currency,omitempty" gorm:"type:varchar(3)"`
	SellerId        string           `json:"seller_id,omitempty" gorm:"type:varchar(100);index;"`
	InStock         bool             `json:"in_stock" gorm:"in_stock,omitempty"`
	DeliveryOptions []DeliveryOption `gorm:"many2many:delivery_products;" json:"delivery_options,omitempty"`
}

type DeliveryOption struct {
	gorm.Model
	Name     string    `json:"name,omitempty" gorm:"type:varchar(100)"`
	Price    float32   `json:"price,omitempty"`
	Currency string    `json:"currency,omitempty" gorm:"type:varchar(3)"`
	Products []Product `gorm:"many2many:delivery_products;" json:"products,omitempty"`
}

package repositories

import (
	"app/models"
	"gorm.io/gorm"
)

type DeliveryOptionDb struct {
	Db *gorm.DB
}

type DeliveryOption interface {
	Find() ([]models.DeliveryOption, error)
}

func (h DeliveryOptionDb) Find() ([]models.DeliveryOption, error) {
	var deliveryOptions []models.DeliveryOption

	err := h.Db.Find(&deliveryOptions).Error

	return deliveryOptions, err
}

func NewDeliveryOptionDb(db *gorm.DB) DeliveryOptionDb {
	return DeliveryOptionDb{Db: db}
}

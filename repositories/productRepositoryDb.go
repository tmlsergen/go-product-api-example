package repositories

import (
	"app/dto"
	"app/models"
	"gorm.io/gorm"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

type ProductRepository interface {
	FindById(productId uint) (models.Product, error)
	Create(productDto dto.ProductRequestDto, deliveryOptions []models.DeliveryOption) (models.Product, error)
	Find(orderBy bool) ([]models.Product, error)
	Update(productId uint, productDto dto.ProductRequestDto) (models.Product, error)
}

func (h ProductRepositoryDb) Create(productDto dto.ProductRequestDto, deliveryOptions []models.DeliveryOption) (models.Product, error) {
	var product = models.Product{
		Name:            productDto.Name,
		Currency:        productDto.Currency,
		Price:           productDto.Price,
		Description:     productDto.Description,
		SellerId:        productDto.SellerId,
		InStock:         productDto.InStock,
		DeliveryOptions: deliveryOptions,
	}

	result := h.Db.Create(&product)

	return product, result.Error
}

func (h ProductRepositoryDb) Find(orderBy bool) ([]models.Product, error) {
	var products []models.Product

	result := h.Db

	if orderBy {
		result = result.Order("price DESC")
	}

	err := result.Preload("DeliveryOptions").Find(&products).Error

	return products, err
}

func (h ProductRepositoryDb) Update(productId string, productDto dto.ProductRequestDto) (models.Product, error) {
	var product = models.Product{
		Name:        productDto.Name,
		Currency:    productDto.Currency,
		Price:       productDto.Price,
		Description: productDto.Description,
		SellerId:    productDto.SellerId,
		InStock:     productDto.InStock,
	}

	result := h.Db.Where("id = ?", productId).Updates(&product)

	return product, result.Error
}

func (h ProductRepositoryDb) FindById(productId string) (models.Product, error) {
	var product = models.Product{}

	result := h.Db.Preload("DeliveryOptions").Where("id = ?", productId).First(&product)

	return product, result.Error
}

func NewProductRepositoryDb(db *gorm.DB) ProductRepositoryDb {
	return ProductRepositoryDb{Db: db}
}

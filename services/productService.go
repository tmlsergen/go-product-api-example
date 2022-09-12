package services

import (
	"app/dto"
	"app/models"
	"app/repositories"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"time"
)

type ProductService struct {
	ProductRepository        repositories.ProductRepositoryDb
	DeliveryOptionRepository repositories.DeliveryOptionDb
	RedisClient              *redis.Client
}

type ProductServiceInterface interface {
	CreateProduct(productDto dto.ProductRequestDto) (models.Product, error)
	ListProducts(orderBy bool) ([]models.Product, error)
	UpdateProduct(productId string, productDto dto.ProductRequestDto) (models.Product, error)
	ShowProduct(productId string) (models.Product, error)
}

func (h ProductService) CreateProduct(productDto dto.ProductRequestDto) (models.Product, error) {
	deliveryOptions, err := h.DeliveryOptionRepository.Find()

	if err != nil {
		return models.Product{}, err
	}

	product, err := h.ProductRepository.Create(productDto, deliveryOptions)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (h ProductService) UpdateProduct(productId string, productDto dto.ProductRequestDto) (models.Product, error) {
	product, err := h.ProductRepository.Update(productId, productDto)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (h ProductService) ListProducts(orderBy bool) ([]models.Product, error) {
	ctx := context.Background()

	var products []models.Product

	if orderBy {
		encodedProducts, err := h.RedisClient.Get(ctx, "orderedProducts").Result()

		if err != nil {
			products, err := h.ProductRepository.Find(orderBy)

			if err != nil {
				return nil, err
			}

			cachedProducts, err := json.Marshal(products)

			if err != nil {
				return nil, err
			}

			h.RedisClient.Set(ctx, "orderedProducts", string(cachedProducts), 1*time.Hour)

			return products, nil
		}

		err = json.Unmarshal([]byte(encodedProducts), &products)

		if err != nil {
			return nil, err
		}
	}

	encodedProducts, err := h.RedisClient.Get(ctx, "products").Result()

	if err != nil {
		products, err := h.ProductRepository.Find(orderBy)
		if err != nil {
			return nil, err
		}

		cachedProducts, err := json.Marshal(products)

		if err != nil {
			return nil, err
		}

		h.RedisClient.Set(ctx, "products", string(cachedProducts), 1*time.Hour)

		return products, nil
	}

	err = json.Unmarshal([]byte(encodedProducts), &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (h ProductService) ShowProduct(productId string) (models.Product, error) {
	product, err := h.ProductRepository.FindById(productId)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func NewProductService(productRepo repositories.ProductRepositoryDb, deliveryOptionRepo repositories.DeliveryOptionDb, redisClient *redis.Client) ProductService {
	return ProductService{ProductRepository: productRepo, DeliveryOptionRepository: deliveryOptionRepo, RedisClient: redisClient}
}

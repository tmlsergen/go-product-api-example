package handlers

import (
	"app/dto"
	"app/services"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductService services.ProductService
}

type ProductHandlerInterface interface {
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	ListProducts(c *fiber.Ctx) error
	ShowProduct(c *fiber.Ctx) error
}

func (h ProductHandler) CreateProduct(c *fiber.Ctx) error {
	productDto := new(dto.ProductRequestDto)

	if err := c.BodyParser(productDto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := dto.ValidateProductStruct(*productDto)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	product, err := h.ProductService.CreateProduct(*productDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.ProductResponseDto{
		Name:            product.Name,
		Description:     product.Description,
		Currency:        product.Currency,
		Price:           product.Price,
		SellerId:        product.SellerId,
		InStock:         product.InStock,
		DeliveryOptions: product.DeliveryOptions,
	})
}

func (h ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productDto := new(dto.ProductRequestDto)
	productId := c.Params("id")

	if err := c.BodyParser(productDto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := dto.ValidateProductStruct(*productDto)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	product, err := h.ProductService.UpdateProduct(productId, *productDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.ProductResponseDto{
		Name:            product.Name,
		Description:     product.Description,
		Currency:        product.Currency,
		Price:           product.Price,
		SellerId:        product.SellerId,
		InStock:         product.InStock,
		DeliveryOptions: product.DeliveryOptions,
	})
}

func (h ProductHandler) ListProducts(c *fiber.Ctx) error {
	var productsDto []dto.ProductResponseDto

	products, err := h.ProductService.ListProducts(false)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	for _, product := range products {
		productDto := dto.ProductResponseDto{
			Name:            product.Name,
			Description:     product.Description,
			Currency:        product.Currency,
			Price:           product.Price,
			SellerId:        product.SellerId,
			InStock:         product.InStock,
			DeliveryOptions: product.DeliveryOptions,
		}

		productsDto = append(productsDto, productDto)
	}

	return c.Status(fiber.StatusOK).JSON(productsDto)
}

func (h ProductHandler) ShowProduct(c *fiber.Ctx) error {
	productId := c.Params("id")

	product, err := h.ProductService.ShowProduct(productId)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.ProductResponseDto{
		Name:            product.Name,
		Description:     product.Description,
		Currency:        product.Currency,
		Price:           product.Price,
		SellerId:        product.SellerId,
		InStock:         product.InStock,
		DeliveryOptions: product.DeliveryOptions,
	})
}

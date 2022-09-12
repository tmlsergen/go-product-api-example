package handlers

import (
	"app/dto"
	"app/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	UserService services.UserService
}

type AuthHandlerInterface interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

func (h AuthHandler) Register(c *fiber.Ctx) error {
	userDto := new(dto.UserRequestDto)

	if err := c.BodyParser(userDto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := dto.ValidateUserStruct(*userDto)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	user, err := h.UserService.RegisterUser(*userDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.UserResponseDto{
		Name:  user.Name,
		Email: user.Email,
	})
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	userDto := new(dto.LoginRequestDto)

	if err := c.BodyParser(userDto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := dto.ValidateLoginStruct(*userDto)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	token, err := h.UserService.LoginUser(*userDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.LoginResponseDto{
		Token: token,
	})
}

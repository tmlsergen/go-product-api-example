package middlewares

import (
	"app/auth"
	"github.com/gofiber/fiber/v2"
)

func ValidateAuthRequest(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()

	token := headers["Authorization"]

	err := auth.ValidateToken(token)

	if err != nil {
		return c.Status(401).JSON(err)
	}

	return c.Next()
}

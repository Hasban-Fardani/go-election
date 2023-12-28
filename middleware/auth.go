package middleware

import (
	"go-election/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func Auth() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := jwt.Get(c)
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized, please login first",
			})
		}
		return c.Next()
	}
}

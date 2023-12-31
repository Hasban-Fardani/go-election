package middleware

import (
	"go-election/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func Auth() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token, _ := jwt.Get(c)
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).Redirect("/?message=unauthorized", fiber.StatusFound)
		}
		return c.Next()
	}
}

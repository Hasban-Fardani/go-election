package middleware

import (
	"go-election/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func Admin() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token, claims := jwt.Get(c)
		if token == "" || claims.Role != "admin" {
			return c.Status(fiber.StatusUnauthorized).Redirect("/?message=unauthorized", fiber.StatusFound)
		}
		return c.Next()
	}
}

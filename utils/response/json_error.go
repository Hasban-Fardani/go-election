package response

import "github.com/gofiber/fiber/v2"

func JsonError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": err.Error(),
	})
}

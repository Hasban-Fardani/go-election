package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetCandidates(c *fiber.Ctx) error {
	return c.Render("candidates", fiber.Map{})
}

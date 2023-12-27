package routes

import "github.com/gofiber/fiber/v2"

func AddRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}

package routes

import (
	"go-election/controller"
	"go-election/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
)

func AddRoutes(app *fiber.App) {
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "form:csrf",
		CookieName:     "csrf_token",
		Expiration:     3600,
		CookieSecure:   true,
		CookieHTTPOnly: true,
		KeyGenerator:   utils.UUID,
		ContextKey:     "token",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("message", c.Query("message"))
		return c.Next()
	})

	app.Get("/", controller.Index)
	app.Post("/login", controller.Login)
	app.Post("/logout", controller.Logout)

	app.Get("/election/:id", middleware.Auth(), controller.Election)
	app.Post("/vote/:id", middleware.Auth(), controller.UserVote)

	app.Get("/dashboard", middleware.Admin(), controller.Dashboard)
	app.Get("/admin/elections", middleware.Admin(), controller.AdminElections)

	// app.Get("/api/elections", middleware.Admin(), controller.APIElections)
	app.Get("/api/elections", controller.APIElections)
}

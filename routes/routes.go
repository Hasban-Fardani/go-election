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
	app.Use(compress.New())
	app.Use(logger.New())

	// app.Get("/performence", monitor.New())

	app.Get("/", controller.Index)
	app.Post("/login", controller.Login)
	app.Post("/logout", controller.Logout)

	api := app.Group("/api")
	api.Get("/login/cek", controller.APICheckLogin)
	api.Post("/login", controller.APILogin)
	api.Post("/logout", controller.APILogout)

	elections := api.Group("/elections", middleware.Auth())
	elections.Get("/", controller.GetElections)
	elections.Get("/:id", controller.GetElectionById)
}

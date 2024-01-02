package main

import (
	cfg "go-election/config"
	"go-election/routes"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./web/layout", ".html")
	engine.AddFunc(
		// add unescape function
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	config := fiber.Config{
		Views:     engine,
		AppName:   "go-election",
		Immutable: true,
	}

	app := fiber.New(config)

	// serve ./public as static
	app.Static("/public", "./web/public", fiber.Static{
		Compress: true,
	})

	// add routes
	routes.AddRoutes(app)

	if err := app.Listen(":" + cfg.Data.Port); err != nil {
		panic(err)
	}
}

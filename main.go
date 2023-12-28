package main

import (
	cfg "go-election/config"
	"go-election/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./layout", ".html")
	config := fiber.Config{
		Views:     engine,
		AppName:   "go-election",
		Immutable: true,
	}

	app := fiber.New(config)

	// serve ./public as static
	app.Static("/public", "./public")

	// add routes
	routes.AddRoutes(app)

	if err := app.Listen(":" + cfg.Data.Port); err != nil {
		panic(err)
	}
}

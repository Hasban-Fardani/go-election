package main

import (
	"go-election/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./web", ".html")
	config := fiber.Config{
		Views: engine,
	}

	app := fiber.New(config)

	// serve ./public as static
	app.Static("/public", "./public")

	// add routes
	routes.AddRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

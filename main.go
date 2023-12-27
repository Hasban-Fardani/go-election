package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := fiber.Config{
		ViewsLayout: "./web",
	}

	app := fiber.New(config)

	// serve ./public as static
	app.Static("/public", "./public")

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
